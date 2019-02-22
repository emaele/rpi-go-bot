package private

import (
	"fmt"
	"os/exec"

	gohole "github.com/shuienko/go-pihole"

	"gitlab.com/emaele/rpi-go-bot/commands"
	conf "gitlab.com/emaele/rpi-go-bot/config"
	"gitlab.com/emaele/rpi-go-bot/myip"
	"gitlab.com/emaele/rpi-go-bot/speedtest"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	err error
)

// HandleCommands manages te
func HandleCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message, config conf.Config, ph gohole.PiHConnector) {

	msg := tgbotapi.NewMessage(config.MyID, "")

	switch message.Command() {
	case "start":
		msg.Text = "Hi " + message.From.FirstName + " 👋"
	case "temp":
		msg.Text = fmt.Sprintf("Temperature is %s °C 🔥", commands.GetTemp())
	case "shutdown":
		cmd := exec.Command("shutdown", "-h", "now")
		msg.Text = "Turning off the RPi"
		bot.Send(msg)
		cmd.Run()
	case "reboot":
		cmd := exec.Command("reboot")
		msg.Text = "Rebooting RPi! 🔄"
		bot.Send(msg)
		cmd.Run()
	case "available_space":
		msg.Text, err = commands.AvailableSpace()
		if err != nil {
			msg.Text = err.Error()
		}
	case "speedtest":
		waitMessage := tgbotapi.NewMessage(config.MyID, "Performing a speedtest, please wait... ⏳")
		if sent, err := bot.Send(waitMessage); err == nil {

			ping, down, up := speedtest.Speedtest()

			deleteMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, sent.MessageID)
			bot.Send(deleteMsg)

			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("🕰 Ping: %dms\n\n⬇ Download️: %s\n\n⬆️ Upload: %s", ping, down, up))
			bot.Send(msg)
		}
		return
	case "myip":
		ip, _ := myip.GetMyIP()
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, ip))

	case "pihole":
		if config.Pihole {

			holeArguments := message.CommandArguments()

			switch holeArguments {
			case "enable":
				ph.Enable()
				msg.Text = "Pihole is now enabled ✅"
			case "disable":
				ph.Disable()
				msg.Text = "Pihole is now disabled  🛑"
			default:
				summary := ph.Summary()
				msg.Text = "Pihole is disabled 🛑"
				if summary.Status == "enabled" {
					msg.Text = "Pihole is enabled ✅"
				}
			}
		} else {
			msg.Text = "PiHole disabled in config file"
		}
	default:
		msg.Text = "I don't know that command"
	}
	bot.Send(msg)
}
