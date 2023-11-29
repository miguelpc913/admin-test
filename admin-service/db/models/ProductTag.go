package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductTag struct {
	ProductTagId           int            `json:"productTagId" gorm:"primary_key"`
	Name                   string         `json:"name"`
	Status                 bool           `json:"status"`
	ProductCategoryGeneral bool           `json:"productCategoryGeneral"`
	ProductSet             []Product      `json:"productSet" gorm:"many2many:products_tags;"`
	SalesGroupSet          []SalesGroup   `json:"salesGroupSet" gorm:"many2many:salesgroups_tags;"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at"`
}
