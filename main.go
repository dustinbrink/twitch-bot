package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)


type Config struct {
	IrcUri string
	Nickname string
	OauthToken string
}

func main() {
	log("Starting Twitch Bot")

	config := loadConfig("./config.json")

	irc := IrcClient{nil, config}
	// irc.Props(config)
	irc.Connect()
	irc.Login()
	time.Sleep(time.Second*10)
	irc.Disconnect()
}

func loadConfig(filename string) Config {
	log("loadConfig - loading config file from "+filename)

	config := Config{}
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		exitError("loadConfig - Error opening config file "+filename, err)
	} 

	err = json.Unmarshal(file, &config)
	if err != nil {
		exitError("loadConfig - Error reading config file "+filename, err)
	}

	log("loadConfig - Config loaded for "+ config.Nickname)

  return config
}