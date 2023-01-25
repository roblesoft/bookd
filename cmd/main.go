package main

import (
	"github.com/roblesoft/bookd/pkg/db"
	"github.com/roblesoft/bookd/pkg/models"
	"github.com/roblesoft/bookd/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)
	db.AutoMigrate(&models.Book{})
	db.Create(&models.Book{Author: "Miguel de Cervantes"})
	server := &server.Server{Port: port}
	server.Start()
}