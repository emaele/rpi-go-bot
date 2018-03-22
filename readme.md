<h1>Raspberry Pi Go Bot</h1>

<p>Simple bot written in Golang, I treat it like a school project.<p>

<h3>Config</h3>

Before building it you need to setup your id and token bot. You can get the first one my sending a message to ```@RawDataBot```[^1]

![alt text](img/raw.png =640x480)

```go
var (
	myID     int64 = 000000           
```
TokenBot can be obtained by creating a new bot with ```@BotFather```[^2]

```go
	tokenBot       = "Your token here" 
)
```

<h3>Available commands!</h3>

- ```/temp``` Get current CPU's temperature
- ```/reboot``` Reboot your Raspberry
- ```/available_space``` Get the amount of free GBs on your sd
- ```/speedtest``` Get the result of a speedtest (it requires ```speedtest-cli```)

[^1]: https://t.me/@RawDataBot
[^2]: https://t.me/@BotFather
