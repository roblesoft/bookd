package app

import (
	"log"
	"net"

	http "github.com/roblesoft/bookd/internal/controller/http"
	entity "github.com/roblesoft/bookd/internal/entity"
	"github.com/roblesoft/bookd/internal/usecase"
	repo "github.com/roblesoft/bookd/internal/usecase/repo"
	"github.com/roblesoft/bookd/pkg/db"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func Run() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)
	db.AutoMigrate(&entity.Book{})

	bookRepo := &repo.BookRepository{Db: db}
	service := usecase.NewService(bookRepo)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bookpb.RegisterItemServiceServer(s, &BookServiceServer{repo: repo})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	server := &http.HTTP{Port: port, Service: service}
	server.Start()
}
