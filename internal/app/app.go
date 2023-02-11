package app

import (
	http "github.com/roblesoft/bookd/internal/controller/http"
	entity "github.com/roblesoft/bookd/internal/entity"
	repo "github.com/roblesoft/bookd/internal/usecase/repo"
	"github.com/roblesoft/bookd/pkg/db"
	"github.com/spf13/viper"
)

func Run() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)
	db.AutoMigrate(&entity.Book{})

	repo := &repo.BookRepository{Db: db}

	server := &http.HTTP{Port: port, Repository: repo}
	server.Start()
}
