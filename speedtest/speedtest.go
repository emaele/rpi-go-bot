package speedtest

import (
	"fmt"
	"log"
	"math"

	"github.com/freman/speedtest"
)

// Speedtest performs a speedtest from speedtest.net
func Speedtest() (ping float64, download string, upload string) {

	client := speedtest.NewClient()
	serverList, err := client.GetServerList()
	if err != nil {
		log.Panic(err)
	}

	fastest := serverList.Fastest(5)
	server := fastest[0]

	fmt.Printf("Server: %v", server)

	ping = toFixed(float64(server.TestLatency()/1000000), 2)
	download = fmt.Sprintf("%0.2f mbit/s", server.TestDownload())
	upload = fmt.Sprintf("%0.2f mbit/s", server.TestUpload())

	return

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
