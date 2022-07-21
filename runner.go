package main

import (
	"log"
	"os"

	"example.com/hello/server"
	"example.com/hello/src/handlers"
	"example.com/hello/src/repository"
	"example.com/hello/src/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("configuration impossibe due to: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env file loading failed: %s", err.Error())
	}
	conn, err := repository.NewPsql(repository.Config{
		Host:     viper.GetString("DB.Host"),
		Port:     viper.GetString("DB.Port"),
		Username: viper.GetString("DB.Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("DB.DbName"),
		SSLMode:  viper.GetString("DB.SslMode"),
	})

	if err != nil {
		log.Fatalf("db connection lost %s", err.Error())
	}

	rabbitCongigured, err := repository.NewRabbitConn(repository.RabbitConfig{
		Login:    os.Getenv("RABBITMQ_DEFAULT_USER"),
		Password: os.Getenv("RABBITMQ_DEFAULT_PASS"),
		Port:     os.Getenv("RABBIT_PORT"),
		Host:     os.Getenv("RABBIT_HOST"),
		VHost:    os.Getenv("RABBIT_VHOST"),
	})

	if err != nil {
		log.Fatalf("redis connection lost %s", err.Error())
	}
	repos := repository.NewRepository(conn, rabbitCongigured)
	services := services.NewService(repos)
	server := new(server.Server)
	handlers := handlers.NewHandler(services)
	if err := server.Run(viper.GetString("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("o noooooo. error: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
