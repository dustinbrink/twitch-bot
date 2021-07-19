package main

import (
	"bufio"
	"crypto/tls"
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
	cert, err := tls.LoadX509KeyPair("./publickey.cer", "./private.key")
	if err != nil {
		log("IRC connect - error loading X509 Key Pair")
		log(err.Error())
		return
	}

  config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	i.Conn, err = tls.Dial("tcp", uri, &config)

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

// Parse the irc read line into a Message struct
// Lots more to add here, this is not to spec
// but good enough for my usecase 
func ParseMsg(line string) (m IrcMessage) {
	m = IrcMessage{}
	line = strings.TrimRight(line, "\r\n")

	if line[0] == ':' {
		end := strings.Index(line, " ")
		m.From = line[1:end]
		line = line[end+1:]
	}

	contentStart := strings.Index(line, " :")
	if contentStart > 0 {
		m.Content = line[contentStart+2:]
		line = line[:contentStart]
	}

	fields := strings.Fields(line)
	m.Command = fields[0]
	m.To = fields[1]
	
	// For debug
	// log("Command="+m.Command)
	// log("To="+m.To)
	// log("Content="+m.Content)

	return m
}