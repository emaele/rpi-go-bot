package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	myID      int64   = 000000            // you should replace this with your id
	tokenBot          = "Your token here" // get your token bot from BotFather
	tempLimit float64 = 50                // temperature limit
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
	if err != nil {
		log.Panic(err)
	}

	boot := tgbotapi.NewMessage(myID, "@"+bot.Self.UserName+" is now up! 👌")
	bot.Send(boot)

	go tempAlert(tempLimit, bot)

	for update := range updates {
		go mainBot(bot, update)
	}
}

func mainBot(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	if update.Message.Chat.ID == myID {
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(myID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Hi " + update.Message.From.FirstName + " 👋"
			case "temp":
				cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
				if output, err := getOut(cmd); err == nil {
					log := strings.Split(output, "\n")
					if temp, err := strconv.ParseFloat(log[0], 64); err == nil {
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
					log := strings.Split(output, "\n")

					var down, up string

					for _, element := range log {
						if strings.HasPrefix(element, "Download:") {
							down = element
						} else if strings.HasPrefix(element, "Upload:") {
							up = element
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

func getOut(command *exec.Cmd) (output string, fail error) {
	stdoutStderr, err := command.CombinedOutput()
	if err != nil {
		fail = errors.New("Error")
	}

	output = string(stdoutStderr)

	return output, fail
}

func tempAlert(limit float64, bot *tgbotapi.BotAPI) {

	for {
		msg := tgbotapi.NewMessage(myID, "")
		cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			println("Command error")
		}
		tmp := string(stdoutStderr)
		log := strings.Split(tmp, "\n")
		if temptmp, err := strconv.ParseFloat(log[0], 64); err == nil {

			temp := temptmp / 1000

			if temp >= limit {
				msg.Text = "\t⚠️Attention please⚠️ \n🔥 Your RPi temperature is over " + fmt.Sprint(limit) + "°C 🔥"
				bot.Send(msg)
			}
		} else {
			msg.Text = "Parse error"
			bot.Send(msg)
		}

		time.Sleep(5 * time.Second)
	}
}
