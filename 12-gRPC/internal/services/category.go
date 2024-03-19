package services

import (
	"context"
	"github.com/obrunogonzaga/pos-go-expert/12-gRPC/internal/database"
	"github.com/obrunogonzaga/pos-go-expert/12-gRPC/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.GetName(), in.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating category: %v", err)
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
