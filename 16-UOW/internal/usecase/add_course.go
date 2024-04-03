package usecase

import (
	"context"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/repository"
)

type InputUsecase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUsecase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUsecase(courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepositoryInterface) *AddCourseUsecase {
	return &AddCourseUsecase{
		CourseRepository:   courseRepository,
		CategoryRepository: categoryRepository,
	}
}

func (a *AddCourseUsecase) Execute(ctx context.Context, input InputUsecase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
