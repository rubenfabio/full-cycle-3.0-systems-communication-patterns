package service

import (
	"context"

	"github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns/grpc-module/internal/database"
	"github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns/grpc-module/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.CreateCategory(in.Name, in.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}