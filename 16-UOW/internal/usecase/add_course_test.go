package usecase

import (
	"context"
	"database/sql"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (`id` int PRIMARY KEY AUTO_INCREMENT, `name` varchar(255) NOT NULL)")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (`id` int PRIMARY KEY AUTO_INCREMENT, `name` varchar(255) NOT NULL, `category_id` int NOT NULL)")

	input := InputUsecase{
		CategoryName:     "Golang",
		CourseName:       "Golang for Dummies",
		CourseCategoryID: 1,
	}

	ctx := context.Background()

	useCase := NewAddCourseUsecase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
