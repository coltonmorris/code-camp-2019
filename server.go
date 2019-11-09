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

	r.HandleFunc("/{user}/{service}/playlists", GetPlaylistHandler(api))

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

func GetPlaylistHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		username := vars["user"]
		svc := vars["service"]

		lameUser, ok := api.GetUser(username)
		if !ok {
			fmt.Printf("no user registered for user: %s", username)
		}

		ssvc, okay := lameUser.ServiceAccounts[svc]
		if !okay {
			fmt.Printf("User: %s is not authorized for svc %s", username, svc)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Couldnt get svc")
		}

		_ = ssvc.GetPlaylists()

		api.LoginUser(username)

		w.WriteHeader(http.StatusOK)
		return
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

		redirectUrl := "http://" + r.Host
		fmt.Println("reidrectUrl: ", redirectUrl)

		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}
}

// Uri: /sync/{user}/{playlistName}/{origin_service}/{destination_service}
func SyncHandler(api *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["user"]
		pName := mux.Vars(r)["playlistName"]
		osvc := mux.Vars(r)["origin_service"]
		dsvc := mux.Vars(r)["destination_service"]
		lameUser, ok := api.GetUser(user)
		if !ok {
			fmt.Printf("no user registered for user: %s", user)
		}

		ssvc, okay := lameUser.ServiceAccounts[osvc]
		if !okay {
			fmt.Printf("User: %s is not authorized for svc %s", user, osvc)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Couldnt get svc")
		}
		esvc, okay := lameUser.ServiceAccounts[dsvc]
		if !okay {
			fmt.Printf("User: %s is not authorized for svc %s", user, dsvc)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Couldnt get svc")
		}

		synced, err := TransferPlaylist(pName, pName, ssvc, esvc)
		if err != nil {
			fmt.Print("failed to transfer", err)
		}
		fmt.Println(synced)

	}
}
