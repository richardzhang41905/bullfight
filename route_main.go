package main

import (
	"net/http"

	"github.com/richardzhang41905/bullfight/data"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	games, err := data.Games()
	if err != nil {
		error_message(writer, request, "Cannot get games")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, games, "layout", "public.navbar", "index")
		} else {

			generateHTML(writer, games, "layout", "private.navbar", "index")
		}
	}
}
