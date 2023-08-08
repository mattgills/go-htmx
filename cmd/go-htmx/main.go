package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

type NowHandler struct {
	Now func() time.Time
}

func NewNowHandler(now func() time.Time) NowHandler {
	return NowHandler{Now: now}
}

func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeComponent(nh.Now()).Render(r.Context(), w)
}

func main() {
	component := hello("John")
	header_component := headerTemplate("Paul")
	b := button("Click Me!")
	nf := notFoundComponent()

	http.Handle("/", templ.Handler(component))
	http.Handle("/header", templ.Handler(header_component))
	http.Handle("/button", templ.Handler(b))
	http.Handle("/time", NewNowHandler(time.Now))
	http.Handle("/404", templ.Handler(nf, templ.WithStatus(http.StatusNotFound)))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
