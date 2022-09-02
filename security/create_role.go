package security

import (
	"fmt"
	"org/sonatype/nx/api"
)

type CreateRoleCommand struct {
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *CreateRoleCommand) Execute(args []string) error {
	name := "role"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	if err := createRole(name); err != nil {
		return err
	}

	fmt.Println("Created role", name)
	return nil
}

type rolePayload struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Privileges  []string `json:"privileges"`
}

func createRole(name string) error {
	payload := rolePayload{
		Id:          name,
		Name:        name,
		Description: "",
		Privileges:  []string{"sonatype", "nx-component-upload"},
	}

	return api.Post("v1/security/roles", payload, 200)
}
