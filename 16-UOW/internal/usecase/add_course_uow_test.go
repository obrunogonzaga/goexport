package usecase

import (
	"context"
	"database/sql"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/db"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/repository"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/pkg/uow"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (`id` int PRIMARY KEY AUTO_INCREMENT, `name` varchar(255) NOT NULL)")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (`id` int PRIMARY KEY AUTO_INCREMENT, `name` varchar(255) NOT NULL, `category_id` int NOT NULL, FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`))")

	input := InputUsecase{
		CategoryName:     "Golang",
		CourseName:       "Golang for Dummies",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	uow := uow.NewUow(ctx, dbt)

	uow.Register("category", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("course", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	useCase := NewAddCourseUsecaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
