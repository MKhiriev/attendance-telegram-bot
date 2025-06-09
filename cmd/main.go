package main

import (
	"attendance-telegram-bot/internal/bot/handlers"
	"attendance-telegram-bot/internal/database"
	"attendance-telegram-bot/internal/services"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal("InitConfig(): error loading config.yaml", err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file", err)
	}
	conn := database.NewConnectPostgres()
	repositories := database.NewDatabase(conn)
	service := services.NewService(repositories)
	handler := handlers.NewHandler(service)
	// TODO add bot starter method
	log.Println(handler) // delete me
}

func InitConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
