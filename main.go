package main

import (
	"flag"
	"fmt"
	"github.com/jdiez17/irc-go"
	"os"
	"time"
	"strings"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	configFile := fs.String("config", "", "The config.json")
	fs.Parse(os.Args[1:])

	err := loadConfig(*configFile)
	if err != nil {
		fmt.Println("Error reading the configuration: ", err)
		return
	}

	keyRequests := make(chan string)
	commandStream := make(chan string)

	go Democracy(keyRequests, commandStream)
	go VBAInterface(commandStream)

	conn, err := irc.NewConnection(Config.IRC.Server, int(Config.IRC.Port))
	if err != nil {
		fmt.Println("Error creating connection: ", err)
		return
	}

	conn.Write("PASS " + Config.IRC.Password)
	conn.LogIn(irc.Identity{Nick: Config.IRC.Nick})
	conn.AddHandler(irc.MOTD_END, func(c *irc.Connection, e *irc.Event) {

		for _, channel := range Config.IRC.Channels {
			fmt.Println("Looks good, joining channels")
			c.Join(channel)
		}
	})

	conn.AddHandler(irc.PRIVMSG, func(c *irc.Connection, e *irc.Event) {
		nick := strings.Split(e.Payload["sender"], "!")[0]
		message := e.Payload["message"][:len(e.Payload["message"])-2]
		message := Replace(ToLower(message), " ", "", -1)

		if _, ok := keys[message]; !ok {
			return
		}

		fmt.Println(nick, ":", message)
		keyRequests <- message
	})

	for {
		<-time.After(time.Second)
	}
}
