package main

import (
	"fmt"
	"log"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := page(global.Count, 0)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state
	r.ParseForm()

	// Check to see if the global buttonw as pressed
	if r.Form.Has("global") {
		global.Count++
	}

	// TODO: Update session

	// Display the form
	getHandler(w, r)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	})

	fmt.Println("Listening on localhost:3000")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
