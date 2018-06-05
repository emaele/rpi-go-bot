package main

import (
	"flag"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"gitlab.com/emaele/rpi-go-bot/commands"
	conf "gitlab.com/emaele/rpi-go-bot/config"
	"gitlab.com/emaele/rpi-go-bot/utility"
)

var (
	config         conf.Config
	err            error
	debug          bool
	configFilePath string
)

func main() {

	setCLIParams()

	config, err = conf.ReadConfig(configFilePath)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramTokenBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	boot := tgbotapi.NewMessage(config.MyID, "@"+bot.Self.UserName+" is now up! 👌")
	bot.Send(boot)

	go utility.TempAlert(config.TempLimit, config.MyID, bot)

	for update := range updates {
		if update.Message != nil {
			go mainBot(bot, update.Message, config)
		}
	}
}

func mainBot(bot *tgbotapi.BotAPI, message *tgbotapi.Message, config conf.Config) {

	if message.Chat.ID == config.MyID {
		commands.HandleCommands(bot, message, config)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this bot ⚠️")
		bot.Send(msg)
	}
}

func setCLIParams() {
	flag.BoolVar(&debug, "debug", false, "activate all the debug features")
	flag.StringVar(&configFilePath, "config", "./config.toml", "configuration file path")
	flag.Parse()
}
