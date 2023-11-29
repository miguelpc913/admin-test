package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductInfoType struct {
	ProductInfoTypeId int            `json:"productInfoTypeId" gorm:"primary_key"`
	Id                string         `json:"id"`
	Description       string         `json:"description"`
	DisplayOrder      int            `json:"displayOrder"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}
