package main

import (
	"log"
	"os"
	"os/signal"
	//"github.com/vision8tech/goings/common"
)

var appRoot string
var appServerPort string

func main() {

	// env := common.Env{}

	// Shutting down on interrupt (ctrl/cmd+c) signal.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			log.Printf("Execution was interrupted (signal '%v'). Shutting down ...", sig)
		}
	}()

}
