package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
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

	trx := db.Begin()
	var c Category
	err = trx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", 1).First(&c).Error
	if err != nil {
		trx.Rollback()
		return
	}
	c.Name = "Updated"
	trx.Debug().Save(&c)
	trx.Commit()

	// Create category
	// categoryEletronics := Category{Name: "Electronics"}
	//db.Create(&categoryEletronics)

	// Create another category
	//categoryComputer := Category{Name: "Computer"}
	//db.Create(&categoryComputer)

	// Create product
	//product := Product{Name: "Laptop", Price: 1000, Categories: []Category{categoryEletronics, categoryComputer}}
	//db.Create(&product)

	// Query
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		return
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println(" - ", product.Name)
		}
	}
}
