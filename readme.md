<h1>Raspberry Pi Go Bot</h1>
![alt text1](https://goreportcard.com/badge/gitlab.com/emaele/rpi-go-bot)

<p>RPi bot is a simple bot written in Go to control some aspects of your Raspberry Pi, like cpu temperature</p>
<p>It also checks every X second if CPU temperature goes over the limit of Y°C</p>

<h3>Available commands</h3>

- Get notified on every startup
- ```/temp``` Get current CPU temperature
- ```/reboot``` Reboot your Raspberry
- ```/available_space``` Get the amount of free GBs on your sd
- ```/speedtest``` Get the result of a speedtest (it requires ```speedtest-cli```, you can install it with apt)
- Check constantly for CPU temp and get notified if it reaches a custom value

<h3>Config</h3>

First of all you need to install Telegram Go Apis

```go get github.com/go-telegram-bot-api/telegram-bot-api```

Before building it you need to setup your id,token bot and temperature limit. You can get the first one by sending a message to ```@RawDataBot```[^1]

![alt text](img/raw.png)

```go
var (
	myID     int64 = 000000           
```
TokenBot can be obtained by creating a new bot with ```@BotFather```[^2]

```go
	tokenBot       = "Your token here" 
```
Change the value of tempLimit to set your temperature alert

```go
	tempLimit float64 = 50                // temperature limit
)
```

[^1]: https://t.me/RawDataBot
[^2]: https://t.me/BotFather
