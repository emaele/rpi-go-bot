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
	conf "gitlab.com/emaele/rpi-go-bot/config"
	"github.com/shuienko/go-pihole"
)

var (
	config         conf.Config
	configFilePath = "config.toml"
	err            error
)

func main() {

	config, err = conf.ReadConfig(configFilePath)
	if err != nil {
		log.Panic(err)
	}

	ph := gohole.PiHConnector{
		Host:  config.PiholeHost,
		Token: config.PiholeAPIToken,
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramTokenBot)
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

	boot := tgbotapi.NewMessage(config.MyID, "@"+bot.Self.UserName+" is now up! 👌")
	bot.Send(boot)

	go tempAlert(config.TempLimit, bot)

	for update := range updates {
		if update.Message != nil {
			go mainBot(bot, update, ph)
		}
	}
}

func mainBot(bot *tgbotapi.BotAPI, update tgbotapi.Update, ph gohole.PiHConnector) {

	if update.Message.Chat.ID == config.MyID {
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(config.MyID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Hi " + update.Message.From.FirstName + " 👋"
			case "temp":
				cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
				if stdoutStderr, err := cmd.CombinedOutput(); err == nil {
					tmp := string(stdoutStderr)
					temp := strings.Trim(tmp, "temp='C\n")
					msg.Text = "Temperature is " + temp + "°C 🔥"
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
						msg.Text = "Available space " + fmt.Sprint(value/1000000) + "GB 💾"
					} else {
						msg.Text = "Error"
					}
				} else {
					msg.Text = "Error"
				}
			case "speedtest":
				wait := tgbotapi.NewMessage(config.MyID, "Performing a speedtest, please wait... ⏳")
				bot.Send(wait)
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
			case "pihole":
				holeSwitch := update.Message.CommandArguments()

				switch holeSwitch {
				case "status":
					summary := ph.Summary()
					if summary.Status == "enabled" {
						msg.Text = "Pihole is enabled ✅"
					} else {
						msg.Text = "Pihole is disabled 🛑"
					}
				case "enable":
					ph.Enable()
					msg.Text = "Pihole is now enabled ✅"
				case "disable":
					ph.Disable()
					msg.Text = "Pihole is now disabled  🛑"
				default:
					msg.Text = "Argument not recognized, use status/enable/disable"
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
		msg := tgbotapi.NewMessage(config.MyID, "")
		cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
		if stdoutStderr, err := cmd.CombinedOutput(); err == nil {
			tmp := string(stdoutStderr)
			tmptemp := strings.Trim(tmp, "temp='C\n")
			if temp, err := strconv.ParseFloat(tmptemp, 64); err == nil {
				if temp >= limit {
					msg.Text = "\t⚠️Attention please⚠️ \n🔥 Your RPi temperature is over " + fmt.Sprint(limit) + "°C 🔥"
					bot.Send(msg)
				}
			}
		} else {
			msg.Text = "Parse error"
			bot.Send(msg)
		}

		time.Sleep(10 * time.Second)
	}
}
