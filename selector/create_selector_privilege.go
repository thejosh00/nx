package selector

import (
	"fmt"
	"org/sonatype/nx/api"
)

type SelectorCreatePrivilegeCommand struct {
	Expression      string `short:"f" long:"format" default:"maven" description:"format for privilege"`
	ContentSelector string `short:"s" long:"selector-name" required:"true" description:"name of content selector"`
	Positional      struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *SelectorCreatePrivilegeCommand) Execute(args []string) error {
	name := "content-selector-privilege"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	if err := createSelectorPrivilege(name, cmd.ContentSelector); err != nil {
		return err
	}

	fmt.Println("Created content selector privilege", name)
	return nil
}

type privilegePayload struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Actions         []string `json:"actions"`
	Format          string   `json:"format"`
	Repository      string   `json:"repository"`
	ContentSelector string   `json:"contentSelector"`
}

func createSelectorPrivilege(name string, contentSelector string) error {
	payload := privilegePayload{
		Name:            name,
		Description:     "",
		Actions:         []string{"ALL"},
		Format:          "*",
		Repository:      "*",
		ContentSelector: contentSelector,
	}

	return api.Post("v1/security/privileges/repository-content-selector", payload, 201)
}
