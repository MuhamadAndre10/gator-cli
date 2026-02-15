package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(state *State, cmd *Command) error {
	// check it args is not empty
	if len(cmd.Args) != 3 {
		return errors.New("Please provide a command")
	}

	// get user
	user, err := state.Queries.GetUserByName(context.Background(), cmd.Args[2])
	if err != nil {
		return err
	}

	if user.Name == "" {
		return errors.New("user not found")
	}

	state.SetUser(user.Name)

	// print success message
	fmt.Println("User logged in successfully")

	return nil
}
