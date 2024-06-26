package main

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	botToken string
	mongoUri string
}

func loadEnv() env {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return env{
		botToken: os.Getenv("BOT_TOKEN"),
		mongoUri: os.Getenv("MONGO_URI"),
	}
}
