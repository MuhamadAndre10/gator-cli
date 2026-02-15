package command

import (
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func FollowHandler(state *State, cmd *Command, user database.User) error {

	if len(cmd.Args) != 3 {
		return errors.New("Please provide a feed name")
	}

	// context with cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate id for the feed user table
	id, _ := uuid.NewV7()

	// get feed from database to get feed id
	feedDB, err := state.Queries.GetFeedByUrl(ctx, cmd.Args[2])
	if err != nil {
		return err
	}

	uFeed, err := state.Queries.CreateFeedUser(ctx, database.CreateFeedUserParams{
		ID:     id,
		UserID: user.ID,
		FeedID: feedDB.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s followed %s\n", uFeed.UserName, uFeed.FeedName)

	return nil
}
