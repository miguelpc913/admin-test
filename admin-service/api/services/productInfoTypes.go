package services

import (
	"admin-v2/api/helpers"
	"admin-v2/db/models"
	"net/http"
)

func (serviceManager *ServiceManager) GetProductInfoType(w http.ResponseWriter, r *http.Request) {
	var productInfoTypes []models.ProductInfoType
	err := serviceManager.db.Find(&productInfoTypes).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error in db query"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, productInfoTypes)
}
