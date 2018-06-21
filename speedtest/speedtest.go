package speedtest

import (
	"fmt"
	"log"

	"github.com/freman/speedtest"
)

// Speedtest performs a speedtest from speedtest.net
func Speedtest() (ping float32, download string, upload string) {

	client := speedtest.NewClient()
	serverList, err := client.GetServerList()
	if err != nil {
		log.Panic(err)
	}

	fastest := serverList.Fastest(2)
	server := fastest[0]

	ping = float32(server.TestLatency())
	download = fmt.Sprintf("%0.2fmbit/s\n", server.TestDownload())
	upload = fmt.Sprintf("%0.2fmbit/s\n", server.TestUpload())

	return

}
