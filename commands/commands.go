package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	conf "gitlab.com/emaele/rpi-go-bot/config"
	"gitlab.com/emaele/rpi-go-bot/speedtest"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shuienko/go-pihole"

	"gitlab.com/emaele/rpi-go-bot/utility"
)

// HandleCommands manages te
func HandleCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message, config conf.Config) {

	msg := tgbotapi.NewMessage(config.MyID, "")

	switch message.Command() {
	case "start":
		msg.Text = "Hi " + message.From.FirstName + " 👋"
	case "temp":
		msg.Text = fmt.Sprintf("Temperature is %s °C 🔥", utility.GetTemp())
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

		/*cmd := exec.Command("speedtest-cli")
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
		} */
		ping, down, up := speedtest.Speedtest()
		msg.Text = fmt.Sprintf("🕰 Ping: %fms\n⬇ Download️: %s\n⬆️ Upload: %s", ping, down, up)
		//msg.Text = "🕰 Ping: "+ ping + "⬇ Download️: " + down + "\n" + "⬆️ Upload: " + up
	case "pihole":

		if config.Pihole {
			ph := gohole.PiHConnector{
				Host:  config.PiholeHost,
				Token: config.PiholeAPIToken,
			}

			holeArguments := message.CommandArguments()

			switch holeArguments {
			case "status":
				summary := ph.Summary()
				msg.Text = "Pihole is disabled 🛑"
				if summary.Status == "enabled" {
					msg.Text = "Pihole is enabled ✅"
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
		} else {
			msg.Text = "PiHole disabled in config file"
		}
	default:
		msg.Text = "I don't know that command"
	}
	bot.Send(msg)
}

func getOut(command *exec.Cmd) (output string, fail error) {
	stdoutStderr, err := command.CombinedOutput()
	if err != nil {
		fail = errors.New("Error")
	}

	output = string(stdoutStderr)

	return output, fail
}
