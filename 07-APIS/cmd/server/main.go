package main

import (
	"fmt"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/configs"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/database"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
