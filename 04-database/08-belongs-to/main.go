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
	//category := Category{Name: "Laptop"}
	//db.Create(&category)

	// Create product
	//product := Product{Name: "Macbook Pro", Price: 2000, CategoryID: category.ID}
	//db.Create(&product)

	// Find all
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		println(product.Name, product.Category.Name)
	}

}
