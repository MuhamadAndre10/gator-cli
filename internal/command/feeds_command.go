package command

import (
	"context"
	"errors"
	"fmt"
)

func ShowAllFeedsHandler(state *State, cmd *Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("Please provide a feed name")
	}

	fmt.Println("Showing all feeds")

	ctx := context.Background()

	feeds, err := state.Queries.GetUserFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\nURL: %s\nUser: %s\n\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}

	return nil
}
