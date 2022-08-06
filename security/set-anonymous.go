package security

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/util"
)

type SetAnonymousCommand struct {
	Verbose    bool `short:"v" long:"verbose" description:"log verbose debug information"`
	Positional struct {
		Value bool `positional-arg-name:"value"`
	} `positional-args:"yes"`
}

func (cmd *SetAnonymousCommand) Execute(args []string) error {
	if !cmd.Verbose {
		util.StopLogging()
	}

	err := setAnonymous(cmd.Positional.Value)
	if err != nil {
		return err
	}
	fmt.Println("Anonymous set to", cmd.Positional.Value)
	return nil
}

type payload struct {
	Enabled bool   `json:"enabled"`
	UserId  string `json:"userId"`
}

func setAnonymous(value bool) error {
	payload := payload{
		Enabled: value,
		UserId:  "anonymous",
	}

	return api.Put("beta/security/anonymous", payload, 200)
}
