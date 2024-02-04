package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema - Create table
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	// Create
	// db.Create(&Product{Name: "Laptop", Price: 1000})

	// Update
	// db.Model(&Product{}).Where("name = ?", "Laptop").Update("Price", 2000)

	// Delete
	// db.Where("name = ?", "Laptop").Delete(&Product{})

	// List all
	var products []Product
	db.Find(&products)
	for _, product := range products {
		println(product.ID, product.Name, product.Price)
	}

}
