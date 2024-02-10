package main

import "os"

func exitCommand(c *config, args ...string) error {
	os.Exit(0)
	return nil
}
