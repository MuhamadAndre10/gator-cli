package main

import (
	"blog_aggregator/internal/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"blog_aggregator/internal/middleware"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	// create command
	commands := command.NewCommands()

	// add command
	commands.AddCommand("login", command.HandlerLogin)
	commands.AddCommand("register", command.HandlerRegister)
	commands.AddCommand("reset", command.ResetHandler)
	commands.AddCommand("users", command.ListUsersHandler)
	commands.AddCommand("agg", command.HandlerAgg)
	commands.AddCommand("addfeed", middleware.LoggedInMiddleware(command.AddFeedHandler))
	commands.AddCommand("follow", middleware.LoggedInMiddleware(command.FollowHandler))
	commands.AddCommand("following", middleware.LoggedInMiddleware(command.FollowingHandler))
	commands.AddCommand("unfollow", middleware.LoggedInMiddleware(command.UnfollowFeedUserHandler))
	commands.AddCommand("feeds", command.ShowAllFeedsHandler)
	commands.AddCommand("browse", middleware.LoggedInMiddleware(command.BrowseHandler))

	// read config
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if db url is empty, set it from env
	if config.DB_URL == "" {
		config.SetDBUrl("postgres://postgres:postgres@localhost:5433/gator?sslmode=disable")
	}

	db, err := OpenDB(*config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	// create state
	state := &command.State{
		Config:  config,
		Queries: dbQueries,
	}

	// get command from user
	arg := os.Args

	// create command
	cmd := &command.Command{
		Name: arg[1],
		Args: arg,
	}

	// run command
	err = commands.Run(state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func OpenDB(cfg config.Config) (*sql.DB, error) {
	return sql.Open("postgres", cfg.DB_URL)
}
