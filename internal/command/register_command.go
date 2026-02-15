package command

import (
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func HandlerRegister(state *State, cmd *Command) error {
	// check it args is not empty
	if len(cmd.Args) != 3 {
		return errors.New("Please provide a username")
	}

	// insert user
	user, err := state.Queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: cmd.Args[2],
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("user already exists")
		}
		return err
	}

	// set user
	state.SetUser(user.Name)

	// print success message
	fmt.Println("User registered successfully")

	return nil
}
