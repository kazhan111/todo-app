package main

import (
	"github.com/kazhan111/todo-app"
	"github.com/kazhan111/todo-app/pkg/handler"
	"github.com/kazhan111/todo-app/pkg/repository"
	"github.com/kazhan111/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	svr := new(todo.Server)
	if err := svr.Run(viper.GetString("8000"), handlers.InitRouter()); err != nil {
		log.Fatalf("an error occurred when starting the http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
