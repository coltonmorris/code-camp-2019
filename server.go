package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RunHttpServer(api *API) {
	// heroku creates this env var automagically. Default use PORT env var, else, use cli flag.
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("p", "8080", "port to serve on"))
	}
	flag.Parse()

	rootDir := "./build"

	r := mux.NewRouter()

	r.HandleFunc("/login/{user}", LoginHandler(api))

	// This is called after a user has logged in and recieved a valid username. We do this to kick off the authentication process for a service. It returns a url
	r.HandleFunc("/register/{user}/{service}", RegisterHandler(api))

	// TODO work down from here enabling each endpoint
	// TODO /callback/{service}/  not working
	r.HandleFunc("/callback/{service}", AuthCallbackHandler)

	// The function for syncing a users playlist
	r.HandleFunc("/sync/{user}/{playlistName}/{service_a}/{service_b}", AuthCallbackHandler)

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

func LoginHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		username := vars["user"]

		api.LoginUser(username)

		w.WriteHeader(http.StatusOK)
		return
	}
}

func RegisterHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")

		vars := mux.Vars(r)

		username := vars["user"]
		// log in the user just in case
		api.LoginUser(username)

		service := vars["service"]

		url := ""
		var err error

		switch service {
		case "spotify":
			url = api.RegisterSpotify(username)
		case "youtube":
			url, err = api.RegisterYoutube(username)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, "ERROR: encountered an error from RegisterYoutube")
				return
			}
		default:
			fmt.Println("unknown service")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "ERROR: invalid service")
			return
		}

		type urlResponse struct {
			url string
		}
		io.WriteString(w, url)
		return
	}
}
