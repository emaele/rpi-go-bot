package utility

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// TempAlert monitors CPU temperature every 10 seconds and sends a message if it's over the limit
func TempAlert(limit float64, myID int64, bot *tgbotapi.BotAPI) {

	for {
		actualTemp, _ := strconv.ParseFloat(GetTemp(), 64)
		if actualTemp >= limit {
			msg := tgbotapi.NewMessage(myID, "\t⚠️Attention please⚠️ \n🔥 Your RPi temperature is over "+fmt.Sprint(limit)+"°C 🔥")
			bot.Send(msg)
		}
		time.Sleep(10 * time.Second)
	}
}

// GetTemp gets the actual temperature of your rpi's CPU
func GetTemp() (temp string) {

	cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
	if stdoutStderr, err := cmd.CombinedOutput(); err == nil {
		log := string(stdoutStderr)
		temp = strings.Trim(log, "temp='C\n")
	}
	return
}
