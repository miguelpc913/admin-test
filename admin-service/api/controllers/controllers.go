package controllers

import (
	"admin-v2/api/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	sm := services.NewServiceManager(db)

	r.Get("/productTags", sm.GetProductTags)

	r.Route("/productInfos", func(r chi.Router) {
		r.Get("/", sm.GetProductInfos)
		r.Get("/{id}", sm.GetProductInfoById)
		r.Put("/{id}/identity", sm.PutProductInfoIdentity)
		r.Put("/{id}/configurations", sm.PutProductInfoSettings)
		r.Put("/{id}/salesGroups", sm.PutProductInfoSalesGroups)
		r.Put("/{id}/products", sm.PutProductInfoProducts)
		r.Put("/{id}/venues", sm.PutProductInfoVenue)
		r.Post("/", sm.PostProductInfos)
		r.Get("/types", sm.GetProductInfoType)
		r.Put("/order", sm.PutOrderProductInfos)
	})
	r.Route("/recommendationRules", func(r chi.Router) {
		r.Get("/", sm.GetRecommendationRules)
		r.Post("/", sm.PostRecommendationRule)
		r.Get("/{id}", sm.GetRecommendationRuleById)
		r.Put("/orderPriority", sm.PutOrderRecommendationRules)
		r.Put("/{id}/identity", sm.PutRecommendationRuleIdentity)
		r.Put("/{id}/general", sm.PutRecommendationRuleGeneral)
		r.Put("/{id}/validities", sm.PutRecommendationRuleValidities)
		r.Put("/{id}/displays", sm.PutRecommendationDisplay)
		r.Put("/{id}/salesGroups", sm.PutRecommendationSalesGroups)
		r.Put("/{id}/buyerTypes", sm.PutRecommendationBuyerTypes)
	})
	return r
}
