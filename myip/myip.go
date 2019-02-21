package myip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type myip struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
}

//GetMyIP returns current external ip address
func GetMyIP() (string, error) {

	var resp myip

	body, err := get("https://api.myip.com")
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.IP, err
}

func get(url string) (body []byte, gerr error) {

	var response *http.Response

	response, gerr = http.Get(url) // nolint: gosec
	if gerr != nil {
		return
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			gerr = err
		}
	}()

	body, gerr = ioutil.ReadAll(response.Body)
	if gerr != nil {
		return
	}

	return body, nil
}
