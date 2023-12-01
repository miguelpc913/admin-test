package services

import (
	dtoRr "admin-v2/api/dto/recommendationRules"
	"admin-v2/api/helpers"
	"admin-v2/db/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm/clause"
)

func (serviceManager *ServiceManager) GetRecommendationRules(w http.ResponseWriter, r *http.Request) {
	var recommendations []models.RecommendationRule
	productId := r.URL.Query().Get("productId")
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.Page - 1) * pagination.Limit
	var totalRecords int64
	if productId != "" {
		_ = serviceManager.db.Model(&recommendations).Where("product_id = ?", productId).Count(&totalRecords).Limit(pagination.Limit).Offset(offset).Find(&recommendations)
	} else {
		_ = serviceManager.db.Model(&recommendations).Count(&totalRecords).Limit(pagination.Limit).Offset(offset).Find(&recommendations)
	}
	response["recommendationRules"] = recommendations
	response["limit"] = pagination.Limit
	response["page"] = pagination.Page
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetRecommendationRuleById(w http.ResponseWriter, r *http.Request) {
	var recommendationRule models.RecommendationRule
	id := chi.URLParam(r, "id")
	err := sm.db.Preload(clause.Associations).Find(&recommendationRule, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not recommendation with that id"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, recommendationRule)
}

func (sm *ServiceManager) PutOrderRecommendationRules(w http.ResponseWriter, r *http.Request) {
	var recommendationRules []models.RecommendationRule
	var req []dtoRr.DisplayOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, item := range req {
		err := sm.db.Model(&recommendationRules).Where("recommendation_rule_id = ?", item.RecommendationRuleId).Update("priority", item.Priority).Error
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid recommendation id"})
		}
	}

	helpers.WriteJSON(w, http.StatusOK, recommendationRules)
}

func (sm *ServiceManager) PutRecommendationRuleIdentity(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationIdentity
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Model(recommendation).Where("recommendation_rule_id = ?", id).Select("Name", "Status", "OfferingType").Updates(req).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationRuleGeneral(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationGeneral
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err := sm.db.Find(&recommendation, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not recommendation with that id"})
		return
	}
	baseProduct := models.Product{}
	err = sm.db.First(&baseProduct, "product_id = ?", req.ProductId).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Base product is not valid"})
		return
	}

	offeredProduct := models.Product{}
	err = sm.db.First(&offeredProduct, "product_id = ?", req.OfferedProductId).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Offered product is not valid"})
		return
	}
	recommendationRule := models.RecommendationRule{
		Product:          baseProduct,
		ProductId:        req.ProductId,
		OfferedProduct:   offeredProduct,
		OfferedProductId: req.OfferedProductId,
		DirectAddToCart:  req.DirectAddToCart,
	}
	err = sm.db.Model(recommendation).Where("recommendation_rule_id = ?", id).Select("ProductId", "OfferedProductId", "DirectAddToCart").Updates(recommendationRule).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationRuleValidities(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationValidities
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Find(&recommendation, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not recommendation with that id"})
		return
	}
	startDateTime, err := helpers.ParseDateTime(req.StartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartDatetime is not valid"})
		return
	}
	endDatetime, err := helpers.ParseDateTime(req.EndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndDatetime is not valid"})
		return
	}
	eventStartDatetime, err := helpers.ParseDateTime(req.EventStartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventStartDatetime is not valid"})
		return
	}
	eventEndDatetime, err := helpers.ParseDateTime(req.EventEndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventEndDatetime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.StartTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartTime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.EndTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime is not valid"})
		return
	}
	recommendationRule := models.RecommendationRule{
		StartDatetime:              startDateTime,
		EndDatetime:                endDatetime,
		EventStartDatetime:         eventStartDatetime,
		EventEndDatetime:           eventEndDatetime,
		WeekDay:                    req.WeekDay,
		StartTime:                  req.StartTime,
		EndTime:                    req.EndTime,
		SessionOffsetMinutesBefore: req.SessionOffsetMinutesBefore,
		SessionOffsetMinutesAfter:  req.SessionOffsetMinutesAfter,
	}
	err = sm.db.Model(recommendation).Where("recommendation_rule_id = ?", id).Updates(recommendationRule).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationDisplay(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationDisplay
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Find(&recommendation, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not recommendation with that id"})
		return
	}
	err = sm.db.Model(recommendation).Where("recommendation_rule_id = ?", id).Updates(req).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationSalesGroups(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&recommendation, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation info with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	salesGroups := []models.SalesGroup{}
	for _, id := range req {
		salesGroup := models.SalesGroup{}
		if err := sm.db.First(&salesGroup, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "SalesGroups are not valid"})
			return
		}
		salesGroups = append(salesGroups, salesGroup)
	}

	sm.db.Model(&recommendation).Association("SalesGroupSet").Replace(salesGroups)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutRecommendationBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&recommendation, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation info with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	buyerTypes := []models.BuyerType{}
	for _, id := range req {
		buyerType := models.BuyerType{}
		if err := sm.db.First(&buyerType, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Buyer types are not valid"})
			return
		}
		buyerTypes = append(buyerTypes, buyerType)
	}

	sm.db.Model(&recommendation).Association("BuyerTypeSet").Replace(buyerTypes)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PostRecommendationRule(w http.ResponseWriter, r *http.Request) {
	req := &dtoRr.PostRecommendation{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	//Validate dates
	startDateTime, err := helpers.ParseDateTime(req.StartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartDatetime is not valid"})
		return
	}
	endDatetime, err := helpers.ParseDateTime(req.EndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndDatetime is not valid"})
		return
	}
	eventStartDatetime, err := helpers.ParseDateTime(req.EventStartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventStartDatetime is not valid"})
		return
	}
	eventEndDatetime, err := helpers.ParseDateTime(req.EventEndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventEndDatetime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.StartTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartTime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.EndTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime is not valid"})
		return
	}

	//Manage associations
	salesGroups := []models.SalesGroup{}
	for _, id := range req.SalesGroupSet {
		salesGroup := models.SalesGroup{}
		if err := sm.db.First(&salesGroup, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "SalesGroups are not valid"})
			return
		}
		salesGroups = append(salesGroups, salesGroup)
	}

	buyerTypes := []models.BuyerType{}
	for _, id := range req.BuyerTypeSet {
		buyerType := models.BuyerType{}
		if err := sm.db.First(&buyerType, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
			return
		}
		buyerTypes = append(buyerTypes, buyerType)
	}

	//Find products
	baseProduct := models.Product{}
	err = sm.db.First(&baseProduct, "product_id = ?", req.ProductId).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Base product is not valid"})
		return
	}

	offeredProduct := models.Product{}
	err = sm.db.First(&offeredProduct, "product_id = ?", req.OfferedProductId).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Offered product is not valid"})
		return
	}

	recommendationRule := models.RecommendationRule{
		Status:                     req.Status,
		Name:                       req.Name,
		OfferingType:               req.OfferingType,
		Product:                    baseProduct,
		ProductId:                  req.ProductId,
		OfferedProduct:             offeredProduct,
		OfferedProductId:           req.OfferedProductId,
		DirectAddToCart:            req.DirectAddToCart,
		StartDatetime:              startDateTime,
		EndDatetime:                endDatetime,
		EventStartDatetime:         eventStartDatetime,
		EventEndDatetime:           eventEndDatetime,
		WeekDay:                    req.WeekDay,
		StartTime:                  req.StartTime,
		EndTime:                    req.EndTime,
		SessionOffsetMinutesBefore: req.SessionOffsetMinutesBefore,
		SessionOffsetMinutesAfter:  req.SessionOffsetMinutesAfter,
		SalesGroupSet:              salesGroups,
		BuyerTypeSet:               buyerTypes,
		Title:                      req.Title,
		Body:                       req.Body,
		Footer:                     req.Footer,
	}
	err = sm.db.Create(&recommendationRule).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, recommendationRule)
}
