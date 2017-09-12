package main

import (
	"fmt"
	"os"

	slack "github.com/mnkd/slackposter"
)

// App is the application object.
type App struct {
	Channel     string
	Message     string
	Username    string
	IconEmoji   string
	SlackConfig slack.Config
}

// Run invoke the App.
func (app *App) Run() int {
	poster := slack.NewSlackPoster(app.SlackConfig)

	if len(app.Channel) > 0 {
		poster.Channel = app.Channel
	}

	if len(app.Username) > 0 {
		poster.Username = app.Username
	}

	if len(app.IconEmoji) > 0 {
		poster.IconEmoji = app.IconEmoji
	}

	err := poster.PostMessage(app.Message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}
	return ExitCodeOK
}
