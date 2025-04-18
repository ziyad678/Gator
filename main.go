package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Ziyad678/Gator/internal/config"
	"github.com/Ziyad678/Gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	logFileName := "app.log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	dbConn, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to db %v\n", err)
	}
	dbQueries := database.New(dbConn)
	s := &state{
		config: &conf,
		db:     dbQueries,
	}
	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
