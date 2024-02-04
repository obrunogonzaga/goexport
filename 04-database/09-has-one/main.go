package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema - Create table
	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	if err != nil {
		return
	}

	// Create category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	// Create product
	product := Product{Name: "Laptop", Price: 1000, Category: category}
	db.Create(&product)

	// Create serial number
	serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	db.Create(&serialNumber)

	// Query products
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
