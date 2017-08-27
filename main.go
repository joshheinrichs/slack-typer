package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/nlopes/slack"
)

var TextFrames = [...]string{
	"typing",
	"typing .",
	"typing ..",
	"typing ...",
	"typing  ..",
	"typing   .",
}

func main() {
	var slackToken, user string
	flag.StringVar(&slackToken, "slack-token", "",
		"A Slack API token, generatable at https://api.slack.com/custom-integrations/legacy-tokens")
	flag.StringVar(&user, "user", "",
		"Name of the user to type at.")
	flag.Parse()
	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "flag is required: -slack-token\n")
		flag.Usage()
		os.Exit(2)
	}
	if user == "" {
		fmt.Fprintf(os.Stderr, "flag is required: -user\n")
		flag.Usage()
		os.Exit(2)
	}

	api := slack.New(slackToken)
	users, err := api.GetUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	userID := ""
	for _, u := range users {
		if u.Name == user {
			userID = u.ID
		}
	}
	if userID == "" {
		fmt.Fprintf(os.Stderr, "error: invalid user\n")
		os.Exit(1)
	}

	_, _, channelID, err := api.OpenIMChannel(userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	f := bufio.NewWriter(os.Stdout)
	i := 0
	for {
		rtm.SendMessage(rtm.NewTypingMessage(channelID))
		time.Sleep(time.Second)
		if isatty.IsTerminal(os.Stdout.Fd()) {
			f.WriteString(TextFrames[i] + "\r")
			f.Flush()
			f.WriteString("\033[K")
			i = (i + 1) % len(TextFrames)
		}
	}
}
