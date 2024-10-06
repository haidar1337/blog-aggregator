package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/haidar1337/gator/internal/config"
	"github.com/haidar1337/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	s := state{
		config: &cfg,
		db:     dbQueries,
	}

	commands := commands{
		cmds: make(map[string]func(s *state, cmd command) error, 0),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("users", handlerGetUsers)
	commands.register("reset", reset)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	cmdArgs := make([]string, 0)
	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			cmdArgs = append(cmdArgs, args[i])
		}
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

}
