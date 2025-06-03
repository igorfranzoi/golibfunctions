package repositories

import (
	"math"

	"github.com/igorfranzoi/base-lib-functions/database/models"
	"gorm.io/gorm"
)

// Retorna um Scope do gorm para ser utilizado na execução da query para paginação
func Paginate(modelValue interface{}, objPage *models.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(modelValue).Count(&totalRows)

	objPage.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(objPage.GetLimit())))

	objPage.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(objPage.GetOffSet()).Limit(objPage.Limit).Order(objPage.GetSort())
	}
}
