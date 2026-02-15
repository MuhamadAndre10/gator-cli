package command

import (
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"errors"
)

type State struct {
	*config.Config
	*database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	cmd map[string]func(*State, *Command) error
}

func NewCommands() *Commands {
	return &Commands{
		cmd: make(map[string]func(*State, *Command) error),
	}
}

func (c *Commands) AddCommand(name string, cmd func(*State, *Command) error) {
	c.cmd[name] = cmd
}

func (c *Commands) Run(state *State, cmd *Command) error {
	cmdFunc, ok := c.cmd[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}

	return cmdFunc(state, cmd)
}
