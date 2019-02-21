package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	gohole "github.com/shuienko/go-pihole"
	conf "gitlab.com/emaele/rpi-go-bot/config"
	"gitlab.com/emaele/rpi-go-bot/private"
	"gitlab.com/emaele/rpi-go-bot/utility"
)

var (
	config         conf.Config
	err            error
	debug          bool
	configFilePath string

	ph gohole.PiHConnector
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

	if config.Pihole {
		ph = gohole.PiHConnector{
			Host:  config.PiholeHost,
			Token: config.PiholeAPIToken,
		}
	}

	//Send a message at every boot up
	boot := tgbotapi.NewMessage(config.MyID, "@"+bot.Self.UserName+" is now up! 👌")
	bot.Send(boot)

	//Start temperature monitor
	go utility.TempAlert(config.TempLimit, config.MyID, bot)

	for update := range updates {
		if update.Message != nil {
			go mainBot(bot, update.Message, config)
		}
	}
}

func mainBot(bot *tgbotapi.BotAPI, message *tgbotapi.Message, config conf.Config) {

	action := tgbotapi.NewChatAction(message.Chat.ID, "typing")
	bot.Send(action)

	if message.Chat.ID == config.MyID {
		private.HandleCommands(bot, message, config, ph)
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
