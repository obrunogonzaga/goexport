package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/obrunogonzaga/pos-go-expert/15-SQLC/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend development", Valid: true},
	//})
	//if err != nil {
	//	panic(err)
	//}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "0693cbb1-14eb-49d6-8787-45bc031e8e5d",
		Name:        "Backend Updated",
		Description: sql.NullString{String: "Backend development Updated", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "0693cbb1-14eb-49d6-8787-45bc031e8e5d")

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
