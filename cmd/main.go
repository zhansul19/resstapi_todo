package main

import (
	"os"

	"github.com/joho/godotenv" //loads env vars from a .env file
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"//для логирования
	"github.com/spf13/viper" //application configuration system
	todo "github.com/zhansul19/restapi_todo"
	"github.com/zhansul19/restapi_todo/pcg/handler"
	"github.com/zhansul19/restapi_todo/pcg/repository"
	"github.com/zhansul19/restapi_todo/pcg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))//json формат для колекционирования логов	
	if err := InitConfigs(); err != nil {
		logrus.Fatalf("Error occured setting configuration %s", err.Error())
	}	
	if err:=godotenv.Load(); err != nil {
		logrus.Fatalf("Error occured loadin env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Port:      viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password:  os.Getenv("DB_PASSWORD"),
		DbName:    viper.GetString("db.dbname"),
		SSLmode:   viper.GetString("db.sslmode"),
		Host:      viper.GetString("db.host1"),
	})
	if err != nil {
		logrus.Fatalf("error occured confugiring db: %s", err.Error())
	}
	server := new(todo.Server)
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running Server: %s", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
