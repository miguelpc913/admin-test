package models

import (
	"time"

	"gorm.io/gorm"
)

type SalesGroupType string
type CalendarType string
type PeriodType string
type TicketGroupingType string
type OrderSpecialEventsType string

const (
	WebIndividual         SalesGroupType = "WEB_INDIVIDUAL"
	WebGroups             SalesGroupType = "WEB_GROUPS"
	Seller                SalesGroupType = "SELLER"
	CustomerService       SalesGroupType = "CUSTOMER_SERVICE"
	BoxOffice             SalesGroupType = "BOX_OFFICE"
	SelfService           SalesGroupType = "SELF_SERVICE"
	HTML                  SalesGroupType = "HTML"
	BoxOfficeGroups       SalesGroupType = "BOX_OFFICE_GROUPS"
	CustomerServiceGroups SalesGroupType = "CUSTOMER_SERVICE_GROUPS"
	Numbered              SalesGroupType = "NUMBERED"

	Calendar CalendarType = "calendar"
	Periodic CalendarType = "periodic"

	Day   PeriodType = "day"
	Week  PeriodType = "week"
	Month PeriodType = "month"

	Individual TicketGroupingType = "Individual"
	Collective TicketGroupingType = "Collective"

	DisplayOrder OrderSpecialEventsType = "displayOrder"
	Event        OrderSpecialEventsType = "event"
)

type SalesGroup struct {
	SalesGroupId          int                    `json:"salesGroupId" gorm:"primary_key"`
	ClientId              int                    `json:"clientId"`
	Status                bool                   `json:"status"`
	Name                  string                 `json:"name"`
	DisplayName           string                 `json:"displayName"`
	Slug                  string                 `json:"slug"`
	Type                  SalesGroupType         `json:"type"`
	CalendarStatus        bool                   `json:"calendarStatus"`
	CalendarType          CalendarType           `json:"calendarType"`
	OpeningDate           time.Time              `json:"openingDate"`
	PeriodType            PeriodType             `json:"periodType"`
	NumPeriods            int                    `json:"numPeriods"`
	ClosingDate           time.Time              `json:"closingDate"`
	MinTickets            int                    `json:"minTickets"`
	MaxTickets            int                    `json:"maxTickets"`
	TicketGroupingType    TicketGroupingType     `json:"ticketGroupingType"`
	ReservedEmail         bool                   `json:"reservedEmail"`
	SendEmailReminder     bool                   `json:"sendEmailReminder"`
	SendEmailCancellation bool                   `json:"sendEmailCancellation"`
	ShowPassbookEmail     bool                   `json:"showPassbookEmail"`
	ShowGoogleWalletEmail bool                   `json:"showGoogleWalletEmail"`
	CheckBenefits         bool                   `json:"checkBenefits"`
	LoyaltyProgramId      int                    `json:"loyaltyProgramId"`
	MinNumDays            int                    `json:"minNumDays"`
	PurchaseAfterMinutes  int                    `json:"purchaseAfterMinutes"`
	EditAfterMinutes      int                    `json:"editAfterMinutes"`
	Icon                  string                 `json:"icon"`
	HiddenOnline          bool                   `json:"hiddenOnline"`
	OrderSpecialEvents    OrderSpecialEventsType `json:"orderSpecialEvents"`
	InfoMessage           string                 `json:"infoMessage"`
	DisplayOrder          int                    `json:"displayOrder"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
	DeletedAt             gorm.DeletedAt         `json:"deleted_at"`
	BuyerTypesSet         []BuyerType            `json:"buyerTypesSet" gorm:"many2many:salesgroups_buyertypes;"`
	PaymentMethodsSet     []PaymentMethod        `json:"paymentMethodsSet" gorm:"many2many:salesgroups_paymentmethods;"`
	VerifierSet           []Verifier             `json:"verifiersSet" gorm:"many2many:salesgroups_verifiers;"`
}
