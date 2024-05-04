package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	e := loadEnv()

	dg, err := discordgo.New(e.botToken)
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		fmt.Printf("%s\n", m.Content)
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
