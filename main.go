package main

import (
	"fmt"
	"os"

	"github.com/haidar1337/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	s := state{
		config: &cfg,
	}

	commands := commands{
		cmds: make(map[string]func(s *state, cmd command) error, 0),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	cmdArgs := make([]string, 0)
	if len(args) > 2 {
		cmdArgs = append(cmdArgs, args[2])
	}
	cmdName := args[1]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}
	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
