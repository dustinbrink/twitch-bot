package main

import (
	"bufio"
	"net"
	"net/textproto"
	"strings"
	"time"
)

type IrcClient struct {
	Conn net.Conn
}

type IrcMessage struct {
	To string
	From string
	Content string
	Command string
}

// Connect to the IRC server over tcp
func (i *IrcClient) Connect(uri string) {
	log("IRC connect - connecting to " + uri)
	
	var err error
	i.Conn, err = net.Dial("tcp", uri)

	if err != nil {
		log("IRC connect - error connecting to " + uri)
		log(err.Error())
		return 
	}

	log("IRC connect - connection successful to " + uri + " " + i.Conn.RemoteAddr().String())
}

// Login to IRC with nickname and pass configured in config.json
func (i *IrcClient) Login(nick string, pass string) {
	log("IRC login - logging into IRC as " + nick)
	log("IRC login - oauth token as " + pass)
	i.WriteCommand("PASS "+pass)
	i.WriteCommand("NICK "+nick)
	log("IRC login - log in successful as " + nick)
}

// Join the IRC channel configured in configs.json
func (i *IrcClient) JoinChannel(channelName string) {
	log("IRC join channel - joining channel " + channelName)
	i.WriteCommand("JOIN #"+channelName)
	log("IRC join channel - joined channel " + channelName)
}

// Leave the IRC channel configured in configs.json
func (i *IrcClient) LeaveChannel(channelName string) {
	log("IRC join channel - leaving channel " + channelName)
	i.WriteCommand("PART #"+channelName)
	log("IRC join channel - left channel " + channelName)
}

// Disconnect from the IRC server
func (i *IrcClient) Disconnect() {
	var connIp = i.Conn.RemoteAddr().String()
	i.Conn.Close()
	log("disconnectIRC - connection closed to " + connIp)
}

// Continually read next line from IRC connection and send it to msgHandler()
func (i *IrcClient) WatchChat(msgHandler func(msg IrcMessage)) {
	log("IRC watch chat - watching chat")
	tp := textproto.NewReader((bufio.NewReader(i.Conn)));

	for {
		line, err := tp.ReadLine()

		if err != nil {
			log("IRC watch chat - error reading line")
			log(err.Error())
			i.Disconnect()
			break
		}

		log(line)
		msgHandler(ParseMsg(line));

		time.Sleep(200 * time.Millisecond)
	}
}

// Helper to format IRC commands correctly
func (i *IrcClient) WriteCommand(command string) {
	i.Conn.Write([]byte(command+"\r\n"))
}

// Parse the Irc line into a Message struct
func ParseMsg(line string) (m IrcMessage) {
	m = IrcMessage{}
	line = strings.TrimRight(line, "\r\n")


	if line[0] == ':' {
		end := strings.Index(line, " ")
		m.From = line[1:end]
		line = line[end:]
	}

	return m
}