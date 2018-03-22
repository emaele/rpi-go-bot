package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	myID     int64 = 000000            // you should replace this with your id
	tokenBot       = "Your token here" // get your token bot from BotFather
)

func main() {

	bot, err := tgbotapi.NewBotAPI(tokenBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	boot := tgbotapi.NewMessage(myID, "@"+bot.Self.UserName+" is now up! 👌")
	bot.Send(boot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID == myID {
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(myID, "")
				switch update.Message.Command() {
				case "start":
					msg.Text = "Hi Emanuele 👋"
				case "temp":
					cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
					if output, err := getOut(cmd); err == nil {
						log := strings.Split(output, "\n")
						if temp, err := strconv.ParseFloat(log[0], 32); err == nil {
							temp = temp / 1000
							msg.Text = "Temperature is: " + fmt.Sprint(temp) + "°C 🔥"
						} else {
							msg.Text = "Error"
						}
					} else {
						msg.Text = "Error"
					}
				case "reboot":
					cmd := exec.Command("reboot")
					msg.Text = "Rebooting RPi! 🔄"
					bot.Send(msg)
					cmd.Run()
				case "available_space":
					cmd := exec.Command("df", "--output=avail", "/")
					if output, err := getOut(cmd); err == nil {
						msgSplit := strings.Split(output, "\n")
						if value, err := strconv.Atoi(msgSplit[1]); err == nil {
							msg.Text = "Available space: " + fmt.Sprint(value/1000000) + "GB 💾"
						} else {
							msg.Text = "Error"
						}
					} else {
						msg.Text = "Error"
					}
				case "speedtest":
					cmd := exec.Command("speedtest-cli")
					if output, err := getOut(cmd); err == nil {
						msgSplit := strings.Split(output, "\n")

						var down, up string

						for i := 0; i < len(msgSplit); i++ {
							if strings.HasPrefix(msgSplit[i], "Download:") {
								down = msgSplit[i]
							}
							if strings.HasPrefix(msgSplit[i], "Upload:") {
								up = msgSplit[i]
							}
						}
						msg.Text = "⬇️ " + down + "\n" + "⬆️ " + up
					} else {
						msg.Text = "Error"
					}

				default:
					msg.Text = "I don't know that command"
				}
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "You are not authorized to use this bot ⚠️"
			bot.Send(msg)
		}
	}
}

func getOut(command *exec.Cmd) (output string, fail error) {
	stdoutStderr, err := command.CombinedOutput()
	if err != nil {
		fail = errors.New("Error")
	}

	output = string(stdoutStderr)

	return output, fail
}
