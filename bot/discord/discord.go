package discord

import (
	"bot/counter"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var counterUrl string

func Start(token string, url string) {
	counterUrl = url
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	_, err = dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining discord account details,", err)
		panic(err)
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(checkMessages)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

}

func checkMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := strings.ToLower(m.Content)

	if strings.Contains(message, "!counter") {
		_, champ, _ := strings.Cut(message, "!counter")
		champ = strings.Trim(champ, " ")

		counters, err := counter.FetchCounter(counterUrl, champ)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "champion "+champ+" was not found or errored out")
		} else {
			response := formatCounterMessage(counters)
			_, _ = s.ChannelMessageSend(m.ChannelID, response)
		}

	}
}

func formatCounterMessage(counters counter.Counter) string {
	var builder strings.Builder

	// Append the opening triple backticks for a code block
	builder.WriteString("```\n")

	// Append the champion name
	builder.WriteString(fmt.Sprintf("Counter for %s:\n", counters.Champion))

	// Append the counters as a numbered list
	for i, counterName := range counters.Counters {
		builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, counterName))
	}

	// Append the closing triple backticks for the code block
	builder.WriteString("```")

	return builder.String()
}
