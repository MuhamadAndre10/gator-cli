package command

import (
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AddFeedHandler(state *State, cmd *Command, user database.User) error {

	if len(cmd.Args) != 4 {
		return errors.New("Please provide a feed name, url, and description")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	feedId, err := uuid.NewV7()
	if err != nil {
		return err
	}
	userId := user.ID
	feedName := cmd.Args[2]
	feedUrl := cmd.Args[3]

	_, err = state.Queries.CreateFeed(ctx, database.CreateFeedParams{
		ID:     feedId,
		Name:   feedName,
		Url:    feedUrl,
		UserID: userId,
	})
	if err != nil {
		return err
	}

	_, err = state.Queries.CreateFeedUser(ctx, database.CreateFeedUserParams{
		ID:     feedId,
		UserID: userId,
		FeedID: feedId,
	})
	if err != nil {
		return err
	}

	// user name and feed name
	fmt.Printf("%s added %s successfully\n", user.Name, feedName)

	return nil
}
