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
	IrcChannel string
	IrcRate time.Duration
}

func main() {
	log("Starting Twitch Bot")
	twitch := Twitch{Config: loadConfig("./config.json")}
	twitch.Start()
	time.Sleep(time.Second*300)
	twitch.Stop()
}

// Read local file config.json for settings
// Expected
//    IrcUri string
//	  Nickname string
//	  OauthToken string
//  	IrcChannel string
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