package core

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		} else if pageSize > 100 {
			pageSize = 100
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateG(page int, pageSize int) func(stmt *gorm.Statement) {
	return func(stmt *gorm.Statement) {
		if page < 1 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		} else if pageSize > 100 {
			pageSize = 100
		}
		offset := (page - 1) * pageSize
		stmt.AddClause(clause.Limit{Limit: &pageSize, Offset: offset})
	}
}
