package main

import (
	"os"
	"log"
	"log-detection/application"
)


func main() {
	log.Println(os.Args)
	if len(os.Args) != 2 {
		log.Fatalln("please enter a file config path")
	}
	app := application.NewApplication(os.Args[1])
	app.RunApplication()
}

