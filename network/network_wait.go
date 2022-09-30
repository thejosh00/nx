package network

import (
	"fmt"
	"log"
	"net"
	"org/sonatype/nx/config"
	"os"
	"time"
)

type NetworkWaitCommand struct {
	Host    string `short:"h" long:"host" description:"host to connect to"`
	Port    string `short:"p" long:"port" description:"port to connect to"`
	Timeout int    `short:"t" long:"timeout" description:"seconds to wait before failing"`
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

	wait(host, port, cmd.Timeout)
	return nil
}

func wait(host string, port string, timeout int) {
	fmt.Println("Waiting for Nexus Repository to come online")
	start := time.Now().Unix()
	for {
		log.Println("Connecting to ", host, ":", port)
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)
		if err == nil && conn != nil {
			break
		}
		timeLength := time.Now().Unix() - start
		if timeout != 0 && timeLength > int64(timeout) {
			fmt.Println("Timeout after", timeout, "seconds")
			os.Exit(-1)
		}
		time.Sleep(time.Second)
	}
}
