package models

import (
	"time"

	"gorm.io/gorm"
)

type SalesGroupHtml struct {
	SalesGroupHtmlId int            `json:"salesGroupHtmlId" gorm:"primary_key"`
	Status           bool           `json:"status"`
	Name             string         `json:"name"`
	Icon             string         `json:"icon"`
	Slug             string         `json:"slug"`
	InnerHTML        string         `json:"innerHTML"`
	DisplayOrder     int            `json:"displayOrder"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}
