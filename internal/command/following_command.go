package command

import (
	"blog_aggregator/internal/database"
	"context"
	"fmt"
	"time"
)

func FollowingHandler(state *State, cmd *Command, user database.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	feeds, err := state.Queries.GetFollowingFeeds(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s followed %s\n", feed.UserName, feed.FeedName)
	}

	return nil
}
