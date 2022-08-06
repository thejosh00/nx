package network

import (
	"fmt"
	"log"
	"net"
	"org/sonatype/nx/config"
	"org/sonatype/nx/util"
	"time"
)

type NetworkWaitCommand struct {
	Verbose bool `short:"v" long:"verbose" description:"log verbose debug information"`
}

func (cmd *NetworkWaitCommand) Execute(args []string) error {
	if !cmd.Verbose {
		log.SetOutput(new(util.NoLogger))
	}

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
