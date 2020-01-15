package main

import (
	"discord-bot/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token  string
	Prefix string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	Token = os.Getenv("TOKEN")
	Prefix = os.Getenv("PREFIX")
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(handlers.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
