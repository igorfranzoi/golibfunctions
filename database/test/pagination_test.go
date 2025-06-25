package test

import (
	"context"
	"fmt"
	"testing"

	limit "github.com/igorfranzoi/golibfunctions/config"
	"github.com/igorfranzoi/golibfunctions/database/models"
	"github.com/igorfranzoi/golibfunctions/database/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TestUser is a dummy model for testing the Paginate function.
type TestUser struct {
	ID   uint
	Name string
}

// setupTestDB initializes an in-memory SQLite database and seeds it with data.
func setupTestDB(t *testing.T, recordCount int) *gorm.DB {
	// Using "file::memory:?cache=shared" allows the connection to be shared
	// across different calls in the same process, which is useful for setup.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	// Auto-migrate the schema for our test model.
	err = db.AutoMigrate(&TestUser{})
	require.NoError(t, err)

	// Clean up previous data and seed new data.
	db.Exec("DELETE FROM test_users")
	if recordCount > 0 {
		users := make([]TestUser, recordCount)
		for i := 0; i < recordCount; i++ {
			users[i] = TestUser{Name: fmt.Sprintf("User %d", i+1)}
		}
		err = db.Create(&users).Error
		require.NoError(t, err)
	}

	return db
}

func TestPaginate(t *testing.T) {
	ctx := context.Background()

	// The tests will rely on the default config values (Default: 10, Max: 50)
	// if no environment variables are set.
	totalRecords := 27
	db := setupTestDB(t, totalRecords)

	testCases := []struct {
		name               string
		dbRecordCount      int
		paginationIn       *models.Pagination
		expectedTotalRows  int64
		expectedTotalPages int
		expectedResultLen  int
		expectedFirstID    uint // Assuming default sort is "id desc"
	}{
		{
			name:               "First page with default limit",
			paginationIn:       &models.Pagination{Page: 1, Limit: 0}, // Limit 0 should use default
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 3, // 27 records, 10 per page -> 3 pages
			expectedResultLen:  10,
			expectedFirstID:    27,
		},
		{
			name:               "Second page with specific limit",
			paginationIn:       &models.Pagination{Page: 2, Limit: 10},
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 3,
			expectedResultLen:  10,
			expectedFirstID:    17,
		},
		{
			name:               "Last page (partial)",
			paginationIn:       &models.Pagination{Page: 3, Limit: 10},
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 3,
			expectedResultLen:  7, // Remaining records
			expectedFirstID:    7,
		},
		{
			name:               "Page out of bounds",
			paginationIn:       &models.Pagination{Page: 4, Limit: 10},
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 3,
			expectedResultLen:  0,
		},
		{
			name:               "Custom sort order (asc)",
			paginationIn:       &models.Pagination{Page: 1, Limit: 5, Sort: "id", Order: "asc"},
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 6, // 27 records, 5 per page -> 6 pages
			expectedResultLen:  5,
			expectedFirstID:    1,
		},
		{
			name:               "Limit greater than max limit",
			paginationIn:       &models.Pagination{Page: 1, Limit: 100}, // Should be capped at MaxLimit (50)
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 1, // 27 records, 50 per page -> 1 page
			expectedResultLen:  totalRecords,
			expectedFirstID:    27,
		},
		{
			name:               "Zero page number should default to 1",
			paginationIn:       &models.Pagination{Page: 0, Limit: 10},
			expectedTotalRows:  int64(totalRecords),
			expectedTotalPages: 3,
			expectedResultLen:  10,
			expectedFirstID:    27,
		},
		{
			name:               "Empty table",
			dbRecordCount:      0, // Special case to test with no records
			paginationIn:       &models.Pagination{Page: 1, Limit: 10},
			expectedTotalRows:  0,
			expectedTotalPages: 0,
			expectedResultLen:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: Setup DB for the specific test case if needed
			var testDB *gorm.DB
			if tc.dbRecordCount == 0 && tc.name == "Empty table" {
				testDB = setupTestDB(t, 0)
			} else {
				testDB = db
			}

			var results []TestUser

			// Act
			p, err := repositories.Paginate(ctx, testDB, tc.paginationIn, &results)

			// Assert
			require.NoError(t, err)
			assert.Equal(t, tc.expectedTotalRows, p.TotalRows, "TotalRows mismatch")
			assert.Equal(t, tc.expectedTotalPages, p.TotalPages, "TotalPages mismatch")
			//assert.Equal(t, tc.paginationIn.GetLimit(&config.DefaultConfig), p.Limit, "Limit in struct mismatch")
			assert.Equal(t, tc.paginationIn.GetLimit(&limit.DefaultConfig), p.Limit, "Limit in struct mismatch")
			assert.Len(t, results, tc.expectedResultLen, "Result length mismatch")

			if tc.expectedResultLen > 0 {
				assert.Equal(t, tc.expectedFirstID, results[0].ID, "First element ID mismatch")
			}
		})
	}
}
