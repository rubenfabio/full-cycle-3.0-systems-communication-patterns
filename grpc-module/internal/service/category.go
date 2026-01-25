package service

import (
	"context"
	"io"

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
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.CreateCategory(in.Name, in.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Empty) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAllCategories()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list categories")
	}
	categoryList := &pb.CategoryList{}
	for _, category := range categories {
		categoryList.Categories = append(categoryList.Categories, &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		})
	}
	return categoryList, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindCategoryById(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get category")
	}
	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categoryList := &pb.CategoryList{}
	for {
		categoryRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categoryList)
		}
		if err != nil {
			return err
		}
		categoryResponse, err := c.CategoryDB.CreateCategory(categoryRequest.Name, categoryRequest.Description)
		if err != nil {
			return err
		}
		categoryList.Categories = append(categoryList.Categories, &pb.Category{
			Id: categoryResponse.ID,
			Name: categoryResponse.Name,
			Description: categoryResponse.Description,
		})
	}
}
