package main

import (
	"encoding/json"
	"io/ioutil"
)


type Config struct {
	IrcUri string
	Nickname string
	OauthToken string
	IrcChannel string
	SwansonUri string
}

var SWANSON_URI string

// entry point, start up Twitch Bot with local config.json
func main() {
	log("Starting Twitch Bot")
	config := loadConfig("./config.json")
	SWANSON_URI = config.SwansonUri
	twitch := Twitch{Config: config}
	twitch.Start()
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