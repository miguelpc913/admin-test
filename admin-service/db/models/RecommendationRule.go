package models

import (
	dbhelpers "admin-v2/db/helpers"
	"time"

	"gorm.io/gorm"
)

type RecommendationRule struct {
	RecommendationRuleId       int                    `json:"recommendationRuleId" gorm:"primary_key"`
	Status                     bool                   `json:"status"`
	Name                       string                 `json:"name"`
	OfferingType               string                 `json:"offeringType"`
	ProductId                  int                    `json:"productId"`
	Product                    Product                `json:"-"`
	OfferedProductId           int                    `json:"offeredProductId"`
	OfferedProduct             Product                `json:"-" gorm:"foreignKey:OfferedProductId"`
	DirectAddToCart            bool                   `json:"directAddToCart"`
	StartDatetime              time.Time              `json:"startDatetime"`
	EndDatetime                time.Time              `json:"endDatetime"`
	EventStartDatetime         time.Time              `json:"eventStartDatetime"`
	EventEndDatetime           time.Time              `json:"eventEndDatetime"`
	WeekDay                    dbhelpers.JSONIntSLice `json:"weekDay" gorm:"type:JSON"`
	StartTime                  string                 `json:"startTime"`
	EndTime                    string                 `json:"endTime"`
	SessionOffsetMinutesBefore int                    `json:"sessionOffsetMinutesBefore"`
	SessionOffsetMinutesAfter  int                    `json:"sessionOffsetMinutesAfter"`
	SalesGroupSet              []SalesGroup           `json:"salesGroupSet" gorm:"many2many:recommendation_salesgroups;"`
	BuyerTypeSet               []BuyerType            `json:"buyerTypeSet" gorm:"many2many:recommendation_buyertype;"`
	Title                      string                 `json:"title"`
	Body                       string                 `json:"body"`
	Footer                     string                 `json:"footer"`
	Priority                   int                    `json:"priority"`
	CreatedAt                  time.Time              `json:"-"`
	UpdatedAt                  time.Time              `json:"-"`
	DeletedAt                  gorm.DeletedAt         `json:"-"`
}
