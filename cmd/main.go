package main

import (
	"attendance-telegram-bot/internal/database"
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
	log.Println(repositories) // delete me
}

func InitConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
