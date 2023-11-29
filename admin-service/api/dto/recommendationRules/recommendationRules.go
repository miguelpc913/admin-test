package dtoRr

import "admin-v2/db/models"

type DisplayOrderRequest struct {
	RecommendationRuleId int `json:"recommendationRuleId"`
	Priority             int `json:"priority"`
}

type PostRecommendation struct {
	Status                     bool   `json:"status"`
	Name                       string `json:"name"`
	OfferingType               string `json:"offeringType"`
	ProductId                  int    `json:"productId"`
	OfferedProductId           int    `json:"offeredProductId"`
	DirectAddToCart            bool   `json:"directAddToCart"`
	StartDatetime              string `json:"startDatetime"`
	EndDatetime                string `json:"endDatetime"`
	EventStartDatetime         string `json:"eventStartDatetime"`
	EventEndDatetime           string `json:"eventEndDatetime"`
	WeekDay                    []int  `json:"weekDay"`
	StartTime                  string `json:"startTime"`
	EndTime                    string `json:"endTime"`
	SessionOffsetMinutesBefore int    `json:"sessionOffsetMinutesBefore"`
	SessionOffsetMinutesAfter  int    `json:"sessionOffsetMinutesAfter"`
	SalesGroupSet              []int  `json:"salesGroupSet" `
	BuyerTypeSet               []int  `json:"buyerTypeSet"`
	Title                      string `json:"title"`
	Body                       string `json:"body"`
	Footer                     string `json:"footer"`
}

type GetRecommendation struct {
	RecommendationRuleId       int                 `json:"recommendationRuleId"`
	Status                     bool                `json:"status"`
	Name                       string              `json:"name"`
	OfferingType               string              `json:"offeringType"`
	ProductId                  int                 `json:"productId"`
	OfferedProductId           int                 `json:"offeredProductId"`
	DirectAddToCart            bool                `json:"directAddToCart"`
	StartDatetime              string              `json:"startDatetime"`
	EndDatetime                string              `json:"endDatetime"`
	EventStartDatetime         string              `json:"eventStartDatetime"`
	EventEndDatetime           string              `json:"eventEndDatetime"`
	WeekDay                    []string            `json:"weekDay"`
	StartTime                  string              `json:"startTime"`
	EndTime                    string              `json:"endTime"`
	SessionOffsetMinutesBefore int                 `json:"sessionOffsetMinutesBefore"`
	SessionOffsetMinutesAfter  int                 `json:"sessionOffsetMinutesAfter"`
	SalesGroupSet              []models.SalesGroup `json:"salesGroupSet"`
	BuyerTypeSet               []models.BuyerType  `json:"buyerTypeSet"`
	Title                      string              `json:"title"`
	Body                       string              `json:"body"`
	Footer                     string              `json:"footer"`
	Priority                   int                 `json:"priority"`
}

type PutRecommendationIdentity struct {
	Status       bool   `json:"status"`
	Name         string `json:"name"`
	OfferingType string `json:"offeringType"`
}

type PutRecommendationGeneral struct {
	ProductId        int  `json:"productId"`
	OfferedProductId int  `json:"offeredProductId"`
	DirectAddToCart  bool `json:"directAddToCart"`
}

type PutRecommendationValidities struct {
	StartDatetime              string   `json:"startDatetime"`
	EndDatetime                string   `json:"endDatetime"`
	EventStartDatetime         string   `json:"eventStartDatetime"`
	EventEndDatetime           string   `json:"eventEndDatetime"`
	WeekDay                    []string `json:"weekDay"`
	StartTime                  string   `json:"startTime"`
	EndTime                    string   `json:"endTime"`
	SessionOffsetMinutesBefore int      `json:"sessionOffsetMinutesBefore"`
	SessionOffsetMinutesAfter  int      `json:"sessionOffsetMinutesAfter"`
}

type PutRecommendationDisplay struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Footer string `json:"footer"`
}
