package main

import (
	"net"
)

type IrcClient struct {
	Conn net.Conn
	Config
}

// Connect to the IRC server over tcp
func (i *IrcClient) Connect() {
	log("IRC connect - connecting to " + i.IrcUri)
	
	Conn, err := net.Dial("tcp", i.IrcUri)

	if err != nil {
		log("IRC connect - error connecting to " + i.IrcUri)
		log(err.Error())
		return 
	}

	log("IRC connect - connection successful to " + i.IrcUri)
	i.Conn = Conn
}

func (i *IrcClient) Login() {
	log("IRC login - logging into IRC as " + i.Nickname)
	i.Conn.Write([]byte("PASS "+i.OauthToken+"/r/n"))
	i.Conn.Write([]byte("NICK "+i.Nickname+"/r/n"))

	// if err != nil {
	// 	log(err.Error())
	// 	return 
	// }

	log("IRC login - log in successful as " + i.Nickname)
}

// Disconnect from the IRC server
func (i *IrcClient) Disconnect() {
	i.Conn.Close()
	log("disconnectIRC - connection closed to " + i.IrcUri)
} 