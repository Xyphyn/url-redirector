package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	path = flag.String("yaml", "links.yaml", "The path to the YAML file to parse links for")
	ip   = flag.String("http", "0.0.0.0:6080", "[IP:PORT] the IP and port to run the HTTP server on.")
)

func main() {
	data, err := os.ReadFile(*path)
	if err != nil {
		log.Fatalf("File not found: %s", *path)
	}
	var links []map[string]string
	err = yaml.Unmarshal(data, &links)
	if err != nil {
		log.Fatalf("Failed to parse YAML, make it an array of maps to path, url")
	}
	for _, link := range links {
		fmt.Printf(":: Redirecting %s to %s\n", link["path"], link["url"])
		http.HandleFunc(link["path"], func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, link["url"], http.StatusMovedPermanently)
		})
	}

	fmt.Printf("Listening on %s\n", *ip)
	http.ListenAndServe(*ip, nil)
}
