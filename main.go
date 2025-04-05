package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ziyad678/Gator/internal/config"
)

type state struct {
	config *config.Config
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
	s := &state{
		config: &conf,
	}
	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmd := command{
		name: os.Args[0],
		args: os.Args[2:],
	}
	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
	}

}
