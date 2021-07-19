package main

// hold the configs and irc client
type Twitch struct {
	irc IrcClient
	Config
}

// start the connection with IRC server
func (t *Twitch) Start() {
	t.irc = IrcClient{nil}
	t.irc.Connect(t.Config.IrcUri)
	t.irc.Login(t.Config.Nickname, t.Config.OauthToken)
	t.irc.JoinChannel(t.Config.IrcChannel)
	t.irc.WatchChat(t.handleMsg)
	
}

// Disconnect from IRC server
func (t *Twitch) Stop() {
	t.irc.Disconnect()
}

// Callback to handle received IRC messages
func (t *Twitch) handleMsg(msg IrcMessage) {
	switch msg.Command {
		case "PING":
			t.KeepAlive()
		case "PRIVMSG":
			return
		default:
			log(msg.From)
			return
	}
}

// Send the PONG command back to twitch to keep alive connection
func (t *Twitch) KeepAlive() {
	log("Twitch keep alive - sending pong message")
	t.irc.WriteCommand("PONG :tmi.twitch.tv")
}