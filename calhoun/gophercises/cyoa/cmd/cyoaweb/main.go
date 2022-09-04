package main

import (
	"flag"
	"fmt"
	"github.com/hlthung/golang-learning/calhoun/gophercises/cyoa"
	"log"
	"net/http"
	"os"
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
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	//tpl := template.Must(template.New("").Parse("Hello"))
	//h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))

	h := cyoa.NewHandler(story)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
