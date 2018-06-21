package speedtest

import (
	"fmt"
	"log"

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

	ping = float64(server.TestLatency() / 1000000)
	download = fmt.Sprintf("%0.2f mbit/s", server.TestDownload())
	upload = fmt.Sprintf("%0.2f mbit/s", server.TestUpload())

	return

}
