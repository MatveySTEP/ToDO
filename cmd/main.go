package main

import (
	"ToDoList"
	"ToDoList/package/handler"
	"ToDoList/package/repository"
	"ToDoList/package/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)
	srv := new(ToDoList.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
