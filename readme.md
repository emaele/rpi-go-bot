# Raspberry Pi Go Bot

RPi bot is a simple bot written in Go to control some aspects of your Raspberry Pi, like cpu temperature.
It also checks every 10 second if CPU temperature goes over the limit of 60°C (you can edit this value in the config file).
Be sure to run it as root.
### Available commands and features

- Get notified on every startup
- Check constantly for CPU temp and get notified if it reaches a custom value
- ```/temp``` Get current CPU temperature
- ```/reboot``` Reboot your Raspberry
- ```/available_space``` Get the amount of free GBs on your sd
- ```/speedtest``` Get the result of a speedtest
- ```/myip``` Shows your current external ip address
- ```/pihole``` Get current status of pihole
    - ```/pihole enable/disable``` Enable or disable pihole

### Config

All you need is to create and edit the config file, you can also rename config_example.toml to config.toml to do that.

Telegram Bot Token can be obatined by creating a bot with ```@botfather```[^1] and your id by sending a message to ```@rawdatabot```[^2].

After you done everything, you're ready to build and execute the bot. Type:
```
go build
sudo ./rpi-go-bot
```

[^1]: https://t.me/BotFather
[^2]: https://t.me/RawDataBot

