package main

// hold the configs and irc client
type Twitch struct {
	irc IrcClient
	Config
}

// start the connection with IRC server
func (t *Twitch) Start() {
	t.irc = IrcClient{nil}
	t.irc.Connect(t.Config.IrcUri, t.Config.SslCertPath, t.Config.SslKeyPath)
	t.irc.Login(t.Config.Nickname, t.Config.OauthToken)
	t.irc.JoinChannel(t.Config.IrcChannel)
	t.irc.WatchChat(t.HandleMsg)
	
}

// Disconnect from IRC server
func (t *Twitch) Stop() {
	t.Say("Bye Chat, dustinbrink_bot disconnected.")
	t.irc.Disconnect()
}

// Callback to handle received IRC messages
func (t *Twitch) HandleMsg(msg IrcMessage) {
	switch msg.Command {
		case "PING":
			// Ping-pong keep alive
			t.KeepAlive()
		case "PRIVMSG":
			// investigate message for possible commands
			if msg.Content[0] == '!' {
				t.HandleCommands(msg)
			}
		case "001":
			// connected, say hello to chat
			t.Say("Hello Chat, dustinbrink_bot is connected. Possible commands, !swanson !botstop")
		default:
			//  do nothing
			return
	}
}

// Respond to actual bot commands that start with a !
func (t *Twitch) HandleCommands(msg IrcMessage) {
	switch msg.Content {
	case "!swanson":
		  // fetch & say new swanson quote 
			t.Say("\""+fetchSwansonQuote()+"\" - Ron Swanson")
	case "!botstop":
		  // shut down twitter bot
			t.Stop()
	default:
		  // do nothing
			return
	}
}

func (t *Twitch) Say(s string) {
	msg := "PRIVMSG #" + t.IrcChannel + " :" + s;
	log("Twitch say - " + msg)
	t.irc.WriteCommand(msg)
}



// Send the PONG command back to twitch to keep alive connection
func (t *Twitch) KeepAlive() {
	log("Twitch keep alive - sending pong message")
	t.irc.WriteCommand("PONG :tmi.twitch.tv")
}