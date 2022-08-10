package security

import (
	"fmt"
	"org/sonatype/nx/api"
	"strconv"
)

type SetAnonymousCommand struct {
	Positional struct {
		Value string `positional-arg-name:"value"`
	} `positional-args:"yes"`
}

func (cmd *SetAnonymousCommand) Execute(args []string) error {
	value := true
	if cmd.Positional.Value != "" {
		v, err := strconv.ParseBool(cmd.Positional.Value)
		if err != nil {
			return err
		}
		value = v
	}

	if err := setAnonymous(value); err != nil {
		return err
	}

	fmt.Println("Anonymous set to", value)
	return nil
}

type payload struct {
	Enabled   bool   `json:"enabled"`
	UserId    string `json:"userId"`
	RealmName string `json:"realmName"`
}

func setAnonymous(value bool) error {
	payload := payload{
		Enabled:   value,
		UserId:    "anonymous",
		RealmName: "NexusAuthorizingRealm",
	}

	return api.Put("beta/security/anonymous", payload, 200)
}
