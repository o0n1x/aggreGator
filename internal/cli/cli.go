package cli

import (
	"errors"
	"fmt"

	"github.com/o0n1x/aggreGator/internal/config"
)

type State struct {
	State *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("expected arg 'username' but was not found")
	}
	err := s.State.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("set username as: %s\n", cmd.Args[0])
	return nil
}

func (c *Commands) Run(s *State, cmd Command) error {
	err := c.Commands[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}
