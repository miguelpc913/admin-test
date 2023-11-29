package db

import (
	"admin-v2/db/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	var userDb = os.Getenv("MYSQL_USER")
	var pass = os.Getenv("MYSQL_PASSWORD")
	var url = os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT")
	var dbname = os.Getenv("MYSQL_DATABASE")
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userDb,
		pass,
		url,
		dbname,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.BuyerType{})
	db.AutoMigrate(&models.SalesGroup{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductTag{})
	db.AutoMigrate(&models.SalesGroupHtml{})
	db.AutoMigrate(&models.Verifier{})
	db.AutoMigrate(&models.PaymentMethod{})
	db.AutoMigrate(&models.Venue{})
	db.AutoMigrate(&models.ProductInfo{})
	db.AutoMigrate(&models.ProductInfoType{})
	db.AutoMigrate(&models.RecommendationRule{})
	fmt.Println("Successfully connected!", db)
	return db, nil
}
