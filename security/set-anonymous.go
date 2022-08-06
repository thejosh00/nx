package security

import (
	"errors"
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/util"
	"strconv"
)

type SetAnonymousCommand struct {
	Verbose bool `short:"v" long:"verbose" description:"log verbose debug information"`
}

func (cmd *SetAnonymousCommand) Execute(args []string) error {
	if !cmd.Verbose {
		util.StopLogging()
	}

	value := true
	if len(args) > 0 {
		var err error
		value, err = strconv.ParseBool(args[0])
		if err != nil {
			return errors.New("Please specify a boolean value")
		}
	}

	err := setAnonymous(value)
	if err != nil {
		return err
	}
	fmt.Println("Anonymous set to", value)
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
