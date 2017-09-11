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
$ sendslack -c config.json -m MESSAGE
```

```
$ sendslack -m MESSAGE
```

### Examples

```
$ sendslack 'hello world!
```

## Configuration

### config.json
* Required
* Path `$HOME/.config/sendslack/config.json`

```json
{
  "slack_webhooks": [
    {
      "team": "your-team",
      "channel": "#your-project",
      "username": "Your-Bot-Name",
      "icon_emoji": ":octocat:",
      "webhook_url": "https://hooks.slack.com/services/xxxxx/xxxxx/xxxxxx"
    }
  ]
}
```
