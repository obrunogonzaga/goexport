package usecase

import (
	"context"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/internal/repository"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/16-UOW/pkg/uow"
)

type InputUsecaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUsecaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUsecaseUow(uow uow.UowInterface) *AddCourseUsecaseUow {
	return &AddCourseUsecaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUsecaseUow) Execute(ctx context.Context, input InputUsecase) error {

	return a.Uow.Do(ctx, func(uow *uow.Uow) error {

		category := entity.Category{
			Name: input.CategoryName,
		}
		categoryRepo := a.getCategoryRepository(ctx)
		err := categoryRepo.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		courseRepo := a.getCourseRepository(ctx)
		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUsecaseUow) getCategoryRepository(ctx context.Context) *repository.CategoryRepository {
	repo, err := a.Uow.GetRepository(ctx, "category")
	if err != nil {
		panic(err)
	}
	return repo.(*repository.CategoryRepository)
}

func (a *AddCourseUsecaseUow) getCourseRepository(ctx context.Context) *repository.CourseRepository {
	repo, err := a.Uow.GetRepository(ctx, "course")
	if err != nil {
		panic(err)
	}
	return repo.(*repository.CourseRepository)
}
