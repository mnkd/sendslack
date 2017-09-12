# sendslack
CLI tool for sending a message to Slack incoming webhook.

## Installation

```
$ brew tap mnkd/sendslack
$ brew install sendslack
```

## Usage

```
$ sendslack MESSAGE
```

```
$ sendslack -C config.json -m MESSAGE
```

```
$ sendslack -c CHANNEL -u USERNAME -i :smile: -m MESSAGE
```

### Options
```
-C string
  	/path/to/config.json (default: $HOME/.config/sendslack/config.json)
-c string
  	channel
-i string
  	icon emoji
-m string
  	message
-u string
  	usename
-v	prints current sendslack version
```

### Examples

```
$ sendslack 'hello world!'
```

```
$ sendslack -c general -u Bot -i :smile: -m hello
```

## Configuration

### config.json
* Required
* Path `$HOME/.config/sendslack/config.json`

```json
{
  "slack_webhook": {
    "team": "your-team",
    "channel": "#your-project",
    "username": "Your-Bot-Name",
    "icon_emoji": ":octocat:",
    "webhook_url": "https://hooks.slack.com/services/xxxxx/xxxxx/xxxxxx"
  }
}
```
