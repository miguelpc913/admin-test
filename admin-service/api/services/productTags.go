package services

import (
	"admin-v2/api/helpers"
	"admin-v2/db/models"
	"net/http"
)

func (serviceManager *ServiceManager) GetProductTags(w http.ResponseWriter, r *http.Request) {
	var productTags []models.ProductTag
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.Page - 1) * pagination.Limit
	var totalRecords int64
	_ = serviceManager.db.Model(&productTags).Preload("SalesGroupSet").Count(&totalRecords).Limit(pagination.Limit).Offset(offset).Find(&productTags).Error
	response["productTags"] = productTags
	response["limit"] = pagination.Limit
	response["page"] = pagination.Page
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (serviceManager *ServiceManager) PutOrderProductTags(w http.ResponseWriter, r *http.Request) {
	var productTags []models.ProductTag
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.Page - 1) * pagination.Limit
	var totalRecords int64
	_ = serviceManager.db.Model(&productTags).Preload("SalesGroupSet").Count(&totalRecords).Limit(pagination.Limit).Offset(offset).Find(&productTags).Error
	response["productTags"] = productTags
	response["limit"] = pagination.Limit
	response["page"] = pagination.Page
	helpers.WriteJSON(w, http.StatusOK, response)
}
