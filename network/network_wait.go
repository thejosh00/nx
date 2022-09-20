package network

import (
	"fmt"
	"log"
	"net"
	"org/sonatype/nx/config"
	"time"
)

type NetworkWaitCommand struct {
	Host string `short:"h" long:"host" description:"host to connect to"`
	Port string `short:"p" long:"port" description:"port to connect to"`
}

func (cmd *NetworkWaitCommand) Execute(args []string) error {

	host := config.Host()
	port := config.Port()

	if cmd.Host != "" {
		host = cmd.Host
	}

	if cmd.Port != "" {
		port = cmd.Port
	}

	wait(host, port)
	return nil
}

func wait(host string, port string) {
	fmt.Println("Waiting for Nexus Repository to come online")
	for {
		log.Println("Connecting to ", host, ":", port)
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)
		if err == nil && conn != nil {
			break
		}
		time.Sleep(time.Second)
	}
}
