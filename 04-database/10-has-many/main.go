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
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema - Create table
	err = db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	// Create category
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// categoryProducts := Category{Name: "Products"}
	// db.Create(&categoryProducts)

	// Create product
	// product := Product{Name: "Laptop", Price: 1000, Category: category}
	// db.Create(&product)

	// productMouse := Product{Name: "Mouse", Price: 1000, Category: categoryProducts}
	// db.Create(&productMouse)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name)
		for _, product := range category.Products {
			println(" - "+product.Name, product.Price)
		}
	}
}
