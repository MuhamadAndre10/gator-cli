package command

import (
	"context"
	"fmt"
)

func ListUsersHandler(state *State, cmd *Command) error {

	users, err := state.Queries.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := state.Config.GetUser()

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil

}
