package services

import "gorm.io/gorm"

type ServiceManager struct {
	db *gorm.DB
}

func NewServiceManager(database *gorm.DB) *ServiceManager {
	return &ServiceManager{
		db: database,
	}
}
