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
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("p", "80", "port to serve on"))
	}

	directory := flag.String("d", "./frontend/public/", "the directory of static file to host")
	flag.Parse()

	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*directory))))
	http.Handle("/", http.FileServer(http.Dir("./frontend/public/")))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
