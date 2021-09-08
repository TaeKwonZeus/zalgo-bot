package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/taekwonzeus/botbuilder"
)

func main() {
	session, err := botbuilder.BotData{
		Token:  "",
		Prefix: "z!",
		Commands: []botbuilder.Command{{
			Name:        "zalgo",
			Description: "Turn your text into zalgo text",
			Aliases:     []string{"z"},
			Execute: func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {

			},
		}},
	}.Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Open()

	var msg string

	for msg != "close" {
		fmt.Scanln(&msg)
	}

	session.Close()
}
