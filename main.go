package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db string = "main"

func main() {
	e := loadEnv()

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(e.mongoUri))
	if err != nil {
		panic(err)
	}

	dg, err := discordgo.New("Bot " + e.botToken)
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		newMessage(mongoClient, m)
		fmt.Printf("%s:\t%s\n", m.Author.ID, m.Content)
	})

	if err := dg.Open(); err != nil {
		panic(err)
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
