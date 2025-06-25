package repositories

import (
	"context"

	"github.com/igorfranzoi/golibfunctions/config"
	"github.com/igorfranzoi/golibfunctions/database/models"
	"gorm.io/gorm"
)

// Paginate is a generic function that applies pagination to a GORM query.
// It calculates the total number of records and pages, and then fetches the
// data for the requested page using LIMIT and OFFSET. The results are
// populated into the `out` slice, and the `p` pagination struct is updated
// with the total counts.
func Paginate[T any](ctx context.Context, db *gorm.DB, p *models.Pagination, out *[]T) (*models.Pagination, error) {
	var total int64

	tx := db.WithContext(ctx)

	if err := tx.Model(out).Count(&total).Error; err != nil {
		return nil, err
	}

	limit := p.GetLimit(&config.DefaultConfig)

	p.TotalRows = total
	p.TotalPages = int((total + int64(limit) - 1) / int64(limit))

	err := tx.Limit(limit).
		Offset(p.GetOffset()).
		Order(p.GetSort()).
		Find(out).Error

	return p, err
}
