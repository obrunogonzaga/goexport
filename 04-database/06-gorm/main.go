package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
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
	db.Create(&Product{Name: "Laptop", Price: 1000})

	// Create batch
	db.Create([]Product{
		{Name: "Mouse", Price: 10},
		{Name: "Keyboard", Price: 20},
	})

	// Select one
	var product Product
	//db.First(&product, 1) // find product with integer primary key
	//fmt.Println(product)

	db.First(&product, "name = ?", "Mouse") // find product with name Laptop
	fmt.Println(product)

	// Select all
	var products []Product
	db.Limit(2).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Select with condition, with limit and offset (pagination)
	db.Limit(2).Offset(2).Where("price > ?", 15).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	db.Limit(2).Offset(1).Where("name LIKE ?", "%k%").Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	var p Product
	db.First(&p, 1)
	p.Price = 2000
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2)

	db.Delete(&p2)
}
