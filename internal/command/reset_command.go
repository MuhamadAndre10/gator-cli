package command

import (
	"context"
	"errors"
	"fmt"
)

func ResetHandler(state *State, cmd *Command) error {

	if len(cmd.Args) != 2 {
		return errors.New("Please provide a username")
	}

	err := state.Queries.DeleteAllUser(context.Background())
	if err != nil {
		return err
	}

	state.SetUser("")

	fmt.Println("All users deleted successfully")

	return nil
}
