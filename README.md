# go-tg-bot-config-helper

## Description

Its a simple config helper that can be really handful to use if you start Telegram bots development with Go.
It can inspire you to create some functionality based on fields desribed in config.

## How to use?

Import the lib in your Go project:
```Go
tgconfighelper "github.com/BlinovDev/go-tg-bot-config-helper"
```

And call LoadConfig method:
```Go
config, err := tgconfighelper.LoadConfig()
if err != nil {
	log.Fatalf("Failed to load config: %v", err)
}

bot, err := tgbotapi.NewBotAPI(config.Bot.Token)
if err != nil {
	log.Panic(err)
}
```

## Config example

