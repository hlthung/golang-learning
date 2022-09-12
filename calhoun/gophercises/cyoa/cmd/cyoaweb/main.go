package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hlthung/golang-learning/pkg/utils/httphelper"

	"github.com/hlthung/golang-learning/calhoun/gophercises/cyoa"
)

// go build . && go run cmd/cyoaweb/main.go
func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	// Create our custom CYOA story handler
	//tpl := template.Must(template.New("").Parse("Hello"))
	tpl := template.Must(template.ParseFiles("./template2.html"))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)

	// Note that cyoa.NewHandler can also do the following
	//h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))
	//h := cyoa.NewHandler(story)

	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/story/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here. So this avoided the 404
	mux.Handle("/", cyoa.NewHandler(story))
	// Start the server using our ServeMux
	fmt.Printf("Started server at Port %d\n", *port)
	///log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

	// TODO verify this
	srvCLoser, err := httphelper.ListenAndServeWithClose(fmt.Sprintf(":%d", *port), mux)
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

// Updated chapter parsing function. Technically you don't
// *have* to get the story from the path (it could be a
// header or anything else) but I'm not going to rename this
// since "path" is what we used in the videos.
func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
