package network

import (
	"fmt"
	"net"
	"org/sonatype/nx/config"
	"time"
)

type NetworkWaitCommand struct {
}

func (cmd *NetworkWaitCommand) Execute(args []string) error {
	wait(config.Host(), config.Port())
	return nil
}

func wait(host string, port string) {
	fmt.Println("Waiting for Nexus Repository to come online")
	for {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)
		if err == nil && conn != nil {
			break
		}
		time.Sleep(time.Second)
	}
}
