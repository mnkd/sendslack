package main

import (
	"fmt"
	"os"

	slack "github.com/mnkd/slackposter"
)

// App is the application object.
type App struct {
	Message     string
	SlackConfig slack.Config
}

// Run invoke the App.
func (app *App) Run() int {
	poster := slack.NewSlackPoster(app.SlackConfig)
	err := poster.PostMessage(app.Message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}
	return ExitCodeOK
}
