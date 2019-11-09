package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RunHttpServer() {
	// heroku creates this env var automagically
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("p", "8080", "port to serve on"))
	}
	flag.Parse()

	rootDir := "./build"

	r := mux.NewRouter()
	// TODO /callback/{service}/  not working
	r.HandleFunc("/callback/{service}", AuthCallbackHandler)

	// this root route must come AFTER all other routes to allow other requests through
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(rootDir)))

	http.Handle("/", r)

	log.Printf("Serving %s on HTTP port: %s\n", rootDir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func AuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	// token := ""

	switch vars["service"] {
	case "spotify":
		fmt.Println("Spotify callback hit")
	case "youtube":
		fmt.Println("Youtube code: ", r.FormValue("code"))
	default:
		fmt.Println("Neither spotify or youtube was hit")
	}
}
