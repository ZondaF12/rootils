package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main()  {
	godotenv.Load()

	sess, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, "https://twitter.com") || strings.HasPrefix(m.Content, "https://x.com") {
			s.ChannelMessageDelete(m.ChannelID, m.Reference().MessageID)
			var newUrl string
			if strings.Contains(m.Content, "twitter.com") {
				newUrl = strings.Replace(m.Content, "twitter.com", "fxtwitter.com", 1)
			} else {
				newUrl = strings.Replace(m.Content, "x.com", "fxtwitter.com", 1)
			}

			s.ChannelMessageSend(m.ChannelID, "**Shared By** <@" + m.Author.ID + "> " + newUrl)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}