package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("Token here!")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	test := tgbotapi.NewMessage(8513519, "RPi-go-bot is now up! 👌")
	bot.Send(test)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "temp":
				cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
				stdoutStderr, err := cmd.CombinedOutput()
				if err != nil {
					msg.Text = "Errore comando"
					bot.Send(msg)
				}

				tmp := string(stdoutStderr)
				log := strings.Split(tmp, "\n")
				temp, err := strconv.ParseFloat(log[0], 32)

				if err != nil {
					msg.Text = "Errore parse"
				} else {
					temp = temp / 1000
					msg.Text = "Temperature is: " + fmt.Sprint(temp) + "°C 🔥"
				}
			case "reboot":
				cmd := exec.Command("reboot")
				msg.Text = "Rebooting RPi!"
				bot.Send(msg)
				cmd.Run()

			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}
}
