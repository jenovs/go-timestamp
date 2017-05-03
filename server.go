package main

import (
	"encoding/json"
	"github.com/jenovs/api-timestamp/helpers"
	"log"
	"net/http"
	"os"
)

var dateFormats = []string{
	"Jan 2 2006",
	"2 Jan 2006",
	"2 1 2006",
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Path[1:]

	// if no params serve index file
	if len(param) == 0 {
		http.ServeFile(w, r, "index.html")
		return
	}

	res, _ := helpers.ParseDate(param, dateFormats)
	timestamp, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(timestamp)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func getPort() string {
	if port := os.Getenv("PORT"); len(port) != 0 {
		return ":" + port
	}
	return ":3000"
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}
