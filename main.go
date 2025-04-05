package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ziyad678/Gator/internal/config"
)


func main(){
	logFileName := "app.log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	
	c, err := config.Read()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(c)
	c.SetUser("Test")

}