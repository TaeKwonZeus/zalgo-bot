package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token = ""

func main() {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.AddHandler(onMessage)

	session.Open()
	defer session.Close()

	_, err = fmt.Scanln()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func onMessage(session *discordgo.Session, event *discordgo.MessageCreate) {
	tokens := strings.Fields(event.Content)
	if len(tokens) < 1 {
		return
	}

	if tokens[0] != "z!" {
		return
	}

	if len(tokens) < 2 {
		session.ChannelMessageSend(event.ChannelID, "Argument expected")
		return
	}

	if len(tokens[1]) > 500 {
		session.ChannelMessageSend(event.ChannelID, "Too many characters, max amount is 500")
		return
	}

	session.ChannelMessageSendReply(event.ChannelID, toZalgo(tokens[1]), event.Reference())
}

func toZalgo(in string) string {
	return ""
}
