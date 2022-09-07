package selector

import (
	"fmt"
	"org/sonatype/nx/api"
)

type SelectorCreateCommand struct {
	Expression string `short:"e" long:"expression" default:"path =~ \"/org/|/org/sonatype.*\"" description:"expression for content selector"`
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *SelectorCreateCommand) Execute(args []string) error {
	name := "content-selector"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	if err := createSelector(name, cmd.Expression); err != nil {
		return err
	}

	fmt.Println("Created content selector", name)
	return nil
}

type payload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
}

func createSelector(name string, expression string) error {
	payload := payload{
		Name:        name,
		Description: "",
		Expression:  expression,
	}

	return api.Post("v1/security/content-selectors", payload, 204)
}
