package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hlthung/golang-learning/pkg/utils/httphelper"

	"github.com/hlthung/golang-learning/calhoun/gophercises/urlshort"
)

// Ex2 From https://github.com/gophercises/urlshort
func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")

	//http.ListenAndServe(":8080", yamlHandler)

	// TODO verify this
	srvCLoser, err := httphelper.ListenAndServeWithClose(":8080", yamlHandler)
	if err != nil {
		log.Fatalln("ListenAndServeWithClose Error - ", err)
	}
	// Close HTTP Server
	err = srvCLoser.Close()
	if err != nil {
		log.Fatalln("Server Close Error - ", err)
	}

	log.Println("Server Closed")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
