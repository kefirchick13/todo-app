package main

import (
	"os"

	"github.com/joho/godotenv"
	todo "github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/handler"
	"github.com/kefirchick13/todo-app/pkg/repository"
	"github.com/kefirchick13/todo-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("initConfig error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("initConfig error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error during pushing: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// Подключение к докер контейнеру
// On Windows CMD (not switching to bash):
// docker exec -it <container-id> /bin/sh
// On Windows CMD (after switching to bash):
// docker exec -it <container-id> //bin//sh
// or
// winpty docker exec -it <container-id> //bin//sh
// On Git Bash:
// winpty docker exec -it <container-id> //bin//sh

// Поднятие баз

// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

// Дроп

// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

// запуск контейнера
// docker start 88c9467f71b9

// Просмотр контейнерров(всех)
// docker ps -a
