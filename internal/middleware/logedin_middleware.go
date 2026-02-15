package middleware

import (
	"blog_aggregator/internal/command"
	"blog_aggregator/internal/database"
	"context"
)

// LoggedInMiddleware checks if the user is logged in
func LoggedInMiddleware(handler func(s *command.State, cmd *command.Command, user database.User) error) func(s *command.State, cmd *command.Command) error {
	return func(s *command.State, cmd *command.Command) error {
		user, err := s.Queries.GetUserByName(context.Background(), s.Config.GetUser())
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
