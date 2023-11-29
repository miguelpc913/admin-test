package main

import (
	"admin-v2/api/controllers"
	"admin-v2/db"
	"net/http"
)

func main() {
	database, _ := db.InitDB()
	r := controllers.Init(database)
	http.ListenAndServe(":8080", r)
}
