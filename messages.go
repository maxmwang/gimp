package main

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

const messageCollection string = "messages"

type message struct {
	ServerID  string `bson:"server_id"`
	ChannelID string `bson:"channel_id"`
	MessageID string `bson:"message_id"`
	AuthorID  string `bson:"author_id"`
	Content   string `bson:"content"`
}

func newMessage(mongoClient *mongo.Client, msg *discordgo.MessageCreate) {
	coll := mongoClient.Database(db).Collection(messageCollection)
	m := message{
		ServerID:  msg.GuildID,
		ChannelID: msg.ChannelID,
		MessageID: msg.ID,
		AuthorID:  msg.Author.ID,
		Content:   msg.Content,
	}
	_, err := coll.InsertOne(context.TODO(), m)
	if err != nil {
		fmt.Printf("error adding message to db: %s", err)
	}
}
