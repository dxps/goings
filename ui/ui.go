package main

import (
	"honnef.co/go/js/dom"
)

func run() {

	println("Goings UI starts")

}

func main() {

	var view = dom.GetWindow().Document().(dom.HTMLDocument)
	switch readyState := view.ReadyState(); readyState {
	case "loading":
		view.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		println("[main] Error: Unexpected state: ", readyState)
	}

}
