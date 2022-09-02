package security

import (
	"fmt"
	"org/sonatype/nx/api"
)

type CreateUserCommand struct {
	Positional struct {
		Id string `positional-arg-name:"id"`
	} `positional-args:"yes"`
}

func (cmd *CreateUserCommand) Execute(args []string) error {
	id := "user"
	if cmd.Positional.Id != "" {
		id = cmd.Positional.Id
	}

	if err := createUser(id); err != nil {
		return err
	}

	fmt.Println("Created user", id)
	return nil
}

type userPayload struct {
	Id        string   `json:"userId"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"emailAddress"`
	Password  string   `json:"password"`
	Status    string   `json:"status"`
	Roles     []string `json:"roles"`
}

func createUser(id string) error {
	payload := userPayload{
		Id:        id,
		FirstName: id,
		LastName:  id,
		Email:     id + "@example.com",
		Password:  "password",
		Status:    "active",
		Roles:     []string{"role"},
	}

	return api.Post("v1/security/users", payload, 200)
}
