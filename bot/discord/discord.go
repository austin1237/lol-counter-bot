package discord

import (
	"bot/counter"
	"fmt"
	"strings"
	"time"

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

func daysBetweenNow(unixTimestamp int64) int {
	// Calculate the difference between the two timestamps in nanoseconds.
	duration := time.Since(time.Unix(unixTimestamp, 0))

	// Convert the duration to days by dividing by the number of nanoseconds in a day.
	days := duration.Nanoseconds() / time.Hour.Nanoseconds() / 24

	// Return the number of days.
	return int(days)
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

	builder.WriteString("\n")

	daysSinceUpdate := daysBetweenNow(counters.LastUpdated)

	builder.WriteString("It has been " + fmt.Sprint(daysSinceUpdate) + " days since data was refreshed")

	// Append the closing triple backticks for the code block
	builder.WriteString("```")

	return builder.String()
}
