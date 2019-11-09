/*
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// heroku creates this env var automagically
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("p", "80", "port to serve on"))
	}
	flag.Parse()

	directory := "./build"
	http.Handle("/", http.FileServer(http.Dir(directory)))
	http.Handle("/spotify", handle)

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
