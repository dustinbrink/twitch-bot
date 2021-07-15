package main

import "time"

func main() {
	log("Starting Twitch Bot")

	var irc IrcClient
	irc.setProps("irc.chat.twitch.tv:6697")
	irc.connect()
	time.Sleep(time.Second*10)
	irc.disconnect()
}