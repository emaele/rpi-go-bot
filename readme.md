# Raspberry Pi Go Bot

<p>RPi bot is a simple bot written in Go to control some aspects of your Raspberry Pi, like cpu temperature</p>
<p>It also checks every 10 second if CPU temperature goes over the limit of 60°C (you can edit this value in the config file)</p>
<p>Be sure to run it as root</p>

### Available commands and features

- Get notified on every startup
- ```/temp``` Get current CPU temperature
- ```/reboot``` Reboot your Raspberry
- ```/available_space``` Get the amount of free GBs on your sd
- ```/speedtest``` Get the result of a speedtest (it requires ```speedtest-cli```, you can install it with apt)
- ```/pihole``` status/enable/disable
	- ```status``` get current status of pihole
	. ```enable/disable``` enable or disable pihole
- Check constantly for CPU temp and get notified if it reaches a custom value

### Config

Before building it you need to install go dependencies, run this in a shell:
```
go get github.com/shuienko/go-pihole
go get github.com/BurntSushi/toml
```

Now you need to create and edit the config file, you can also rename config_example.toml to config.toml to do that.

Telegram Bot Token can be obatined by creating a bot with ```@botfather```[^1] and your id by sending a message to ```@rawdatabot```[^2]

After you done everything you're ready to build and execute the bot. Type
```
go build
sudo ./rpi-go-bot
```

[^1]: https://t.me/BotFather
[^2]: https://t.me/RawDataBot

