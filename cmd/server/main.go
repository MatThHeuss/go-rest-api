package main

import (
	"log"
	"net/http"

	"github.com/MatThHeuss/go-rest-api/configs"
	"github.com/MatThHeuss/go-rest-api/internal/entity"
	"github.com/MatThHeuss/go-rest-api/internal/infra/database"
	"github.com/MatThHeuss/go-rest-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
