package main

import (
	"encoding/json"
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

	// if the uri ends in a "/" it will not work
	r.HandleFunc("/callback/{service}", AuthCallbackHandler(api))

	// TODO work down from here enabling each endpoint
	// The function for syncing a users playlist
	r.HandleFunc("/sync/{user}/{playlistName}/{origin_service}/{destination_service}", SyncHandler(api))

	// this root route must come AFTER all other routes to allow other requests through
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(rootDir)))

	http.Handle("/", r)

	log.Printf("Serving %s on HTTP port: %s\n", rootDir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
			fmt.Println("registered youtube: ", url)
		default:
			fmt.Println("unknown service in auth registry")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "ERROR - invalid service: "+service)
			return
		}

		type urlResponse struct {
			url string
		}
		io.WriteString(w, url)
		return
	}
}

func AuthCallbackHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")

		service := mux.Vars(r)["service"]

		switch service {
		case "spotify":
			fmt.Println("Spotify callback hit with request: ", r)
			if err := api.SpotifyAuthCallback(r); err != nil {
				fmt.Println("err from spotify: ", err)
				w.WriteHeader(http.StatusForbidden)
				io.WriteString(w, "Couldn't spotify get token")
				return
			}

		case "youtube":
			fmt.Println("Youtube code: ", r.FormValue("code"))
			if err := api.YoutubeAuthCallback(r); err != nil {
				fmt.Println("err from youtube callback: ", err)
				w.WriteHeader(http.StatusForbidden)
				io.WriteString(w, "Couldn't youtube get token")
				return
			}
		default:
			fmt.Println("unknown service in auth callback")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "ERROR - invalid service: "+service)
			return
		}

		type response struct {
			ok bool
		}
		json.NewEncoder(w).Encode(response{ok: true})
	}
}

// Uri: /sync/{user}/{playlistName}/{origin_service}/{destination_service}
func SyncHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
