package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns/grpc-module/internal/database"
	"github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns/grpc-module/internal/pb"
	"github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns/grpc-module/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}