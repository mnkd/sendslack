package main

import (
	"flag"
	"fmt"
	"os"
)

// Exit codes
const (
	ExitCodeOK int = iota
	ExitCodeError
)

var (
	version  string
	revision string
)

var app = App{}

func init() {
	var path, message, channel, username, iconEmoji string
	var printVersion bool

	flag.BoolVar(&printVersion, "v", false, "prints current sendslack version")
	flag.StringVar(&path, "C", "", "/path/to/config.json (default: $HOME/.config/sendslack/config.json)")
	flag.StringVar(&channel, "c", "", "channel")
	flag.StringVar(&username, "u", "", "usename")
	flag.StringVar(&iconEmoji, "i", "", "icon emoji")
	flag.StringVar(&message, "m", "", "message")
	flag.Parse()

	if printVersion {
		fmt.Fprintln(os.Stdout, "Version:", version)
		fmt.Fprintln(os.Stdout, "Revision:", revision)
		os.Exit(ExitCodeOK)
	}

	config, err := NewConfig(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(ExitCodeError)
	}

	if len(message) == 0 && len(os.Args) == 2 {
		message = os.Args[1]
	}

	if len(message) == 0 {
		fmt.Fprintln(os.Stderr, "Message is empty. To send a message, input a message.")
		os.Exit(ExitCodeError)
	}

	app.Channel = channel
	app.Message = message
	app.IconEmoji = iconEmoji
	app.Username = username
	app.SlackConfig = config.SlackWebhook
}

func main() {
	os.Exit(app.Run())
}
