package command

import (
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"
)

func UnfollowFeedUserHandler(s *State, cmd *Command, user database.User) error {

	if len(cmd.Args) != 3 {
		return errors.New("Please provide a url feed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get feed by url
	feedDB, err := s.Queries.GetFeedByUrl(ctx, cmd.Args[2])
	if err != nil {
		return err
	}

	// delete feed user
	err = s.Queries.DeleteFeedUser(ctx, database.DeleteFeedUserParams{
		UserID: user.ID,
		FeedID: feedDB.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, feedDB.Name)

	return nil
}
