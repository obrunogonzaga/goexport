package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// Create product
	//product := Product{Name: "Laptop", Price: 1000, Category: category}
	//db.Create(&product)

	// Create serial number
	//serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	//db.Create(&serialNumber)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name)
		for _, product := range category.Products {
			println(" - "+product.Name, "Serial Number: "+product.SerialNumber.Number)
		}
	}
}
