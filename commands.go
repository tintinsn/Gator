package main

import "fmt"

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if exists {
		if err := handler(s, cmd); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unknown command %s", cmd.name)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
