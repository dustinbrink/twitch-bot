package main

import (
	"net"
)

type IrcClient struct {
	ServerUri string
	Conn net.Conn
}

// Assign values to the Irc Client (constructor)
func (i *IrcClient) setProps(serverUri string) {
	i.ServerUri = serverUri
}

// Connect to the IRC server over tcp
func (i *IrcClient) connect() {
	log("connectIRC - connecting to " + i.ServerUri)
	
	Conn, err := net.Dial("tcp", i.ServerUri)

	if err != nil {
		log("connectIRC - error connecting to " + i.ServerUri);
		log(err.Error())
		return 
	}

	log("connectIRC - connection successful to " + i.ServerUri);
	i.Conn = Conn
}

// Disconnect from the IRC server
func (i *IrcClient) disconnect() {
	i.Conn.Close()
	log("disconnectIRC - connection closed to " + i.ServerUri);
} 