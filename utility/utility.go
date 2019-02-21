package utility

import (
	"fmt"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/emaele/rpi-go-bot/commands"
)

// TempAlert monitors CPU temperature every 10 seconds and sends a message if it's over the limit
func TempAlert(limit float64, myID int64, bot *tgbotapi.BotAPI) {

	for {
		actualTemp, _ := strconv.ParseFloat(commands.GetTemp(), 64)
		if actualTemp >= limit {
			msg := tgbotapi.NewMessage(myID, "\t⚠️Attention please⚠️ \n🔥 Your RPi temperature is over "+fmt.Sprint(limit)+"°C 🔥")
			bot.Send(msg)
		}
		time.Sleep(10 * time.Second)
	}
}
