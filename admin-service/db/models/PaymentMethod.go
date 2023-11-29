package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ProductId     int            `json:"productId" gorm:"primary_key"`
	Name          string         `json:"name"`
	SalesGroupSet []SalesGroup   `json:"paymentMethodsSet" gorm:"many2many:salesgroups_paymentmethods;"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
