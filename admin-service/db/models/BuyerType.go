package models

import (
	"time"

	"gorm.io/gorm"
)

// Status is a custom type that represents the status of the buyer type.
type Status string

// Enum values for Status
const (
	StatusEnabled  Status = "enabled"
	StatusDisabled Status = "disabled"
)

type BuyerType struct {
	BuyerTypeID           int            `json:"buyerTypeId" gorm:"primary_key"`
	Name                  string         `json:"name"`
	ShortName             string         `json:"shortName"`
	Description           string         `json:"description"`
	Status                bool           `json:"status"`
	DiscountRate          float64        `json:"discountRate"`
	DiscountAmount        float64        `json:"discountAmount"`
	NoHandleAvailability  bool           `json:"noHandleAvailability"`
	Pax                   int            `json:"pax"`
	RequiredPdaColorAlert bool           `json:"requiredPdaColorAlert"`
	AlertColor            string         `json:"alertColor"`
	RequiredPdaSoundAlert bool           `json:"requiredPdaSoundAlert"`
	AlertSound            string         `json:"alertSound"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}
