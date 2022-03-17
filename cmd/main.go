package main

import "log"

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatalln(err.Error())
	}
	app.Start()
}
