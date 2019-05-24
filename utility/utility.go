package utility

import (
	"fmt"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gohole "github.com/shuienko/go-pihole"
	"gitlab.com/emaele/rpi-go-bot/commands"
)

// TempAlert monitors CPU temperature every 10 seconds and sends a message if it's over the limit
func TempAlert(limit float64, myID int64, bot *tgbotapi.BotAPI) {

	for range time.NewTicker(10 * time.Second).C {
		actualTemp, _ := strconv.ParseFloat(commands.GetTemp(), 64)
		if actualTemp >= limit {
			msg := tgbotapi.NewMessage(myID, "\t⚠️Attention please⚠️ \n🔥 Your RPi temperature is over "+fmt.Sprint(limit)+"°C 🔥")
			bot.Send(msg)
		}
	}
}

//GetPiholeSummary returns pihole current status and info as a string
func GetPiholeSummary(ph gohole.PiHConnector) string {
	summary := ph.Summary()

	var status string

	if summary.Status == "enabled" {
		status = "enabled ✅"
	} else {
		status = "disabled 🛑"
	}

	return fmt.Sprintf("PiHole is currently %s\n\nPercentage blocked: %s%%\nDNS Queries: %s\nAds Blocked: %s", status, summary.AdsPercentage, summary.DNSQueries, summary.AdsBlocked)
}
