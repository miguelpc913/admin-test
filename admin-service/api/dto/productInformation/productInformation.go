package dtoPi

type DisplayOrderRequest struct {
	ProductInfoId int `json:"productInfoId"`
	DisplayOrder  int `json:"displayOrder"`
}

type PostProductInfo struct {
	Status        bool     `json:"status"`
	Name          string   `json:"name"`
	InternalName  string   `json:"internalName"`
	Description   string   `json:"description"`
	Icon          string   `json:"icon"`
	InfoType      string   `json:"infoType"`
	CalendarColor string   `json:"calendarColor"`
	Weekdays      []int    `json:"weekDays"`
	SelectedDates []string `json:"selectedDates"`
	SalesGroupSet []int    `json:"salesGroupSet" `
	ProductSet    []int    `json:"productSet"`
	VenueSet      []int    `json:"venueSet"`
}

type GetProductInfo struct {
	Status        bool     `json:"status"`
	Name          string   `json:"name"`
	InternalName  string   `json:"internalName"`
	Description   string   `json:"description"`
	Icon          string   `json:"icon"`
	InfoType      string   `json:"infoType"`
	CalendarColor string   `json:"calendarColor"`
	Weekdays      []int    `json:"weekDays"`
	SelectedDates []string `json:"selectedDates"`
	SalesGroupSet []int    `json:"salesGroupSet" `
	ProductSet    []int    `json:"productSet"`
	VenueSet      []int    `json:"venueSet"`
	DisplayOrder  int      `json:"displayOrder"`
}

type PutProductInfoIdentity struct {
	Status       bool   `json:"status"`
	InternalName string `json:"internalName"`
}

type PutProductInfoSettings struct {
	InfoType      string   `json:"infoType"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Icon          string   `json:"icon"`
	CalendarColor string   `json:"calendarColor"`
	Weekdays      []int    `json:"weekDays"`
	SelectedDates []string `json:"selectedDates"`
}

type PutProductInfoSalesGroups struct {
	SalesGroupSet []int `json:"salesGroupSet" `
}

type PutProductInfoProducts struct {
	ProductSet []int `json:"productSet"`
}
