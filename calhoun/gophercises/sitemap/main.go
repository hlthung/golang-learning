package main

import (
	"flag"
	"fmt"
	"github.com/hlthung/golang-learning/calhoun/gophercises/link"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a diff domain
   5. Find all pages (BFS)
   6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

type empty struct{}

func bfs_empty(urlStr string, maxDepth int) []string {
	seen := make(map[string]empty) // we dont care about the value so put struct{} as this type has the minimum data size / memory
	var q map[string]empty
	nq := map[string]empty{
		urlStr: {},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]empty)
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			// tell me if there's a url key inside the seen map
			if _, ok := seen[url]; ok {
				continue // if seen continue
			}
			seen[url] = empty{} // mark this url as seen
			for _, link := range get(url) {
				nq[link] = empty{}
			}
		}
	}
	ret := make([]string, 0, len(seen)) //create with seen size to optimise
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{}) // we dont care about the value so put struct{} as this type has the minimum data size / memory
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: {},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			// tell me if there's a url key inside the seen map
			if _, ok := seen[url]; ok {
				continue // if seen continue
			}
			seen[url] = struct{}{} // mark this url as seen
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen)) //create with seen size to optimise
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	//return filter(hrefs(resp.Body, base), withSameBase(base))
	return filterSimple(hrefs(resp.Body, base), base)
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)

		}
	}
	return ret
}

// Alternative 1: from the original solution
func filter(links []string, keepLink func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepLink(link) { // if withSameBase then only keep this link
			ret = append(ret, link)
		}
	}
	return ret
}

func withSameBase(baseUrl string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, baseUrl)
	}
}

// Alternative 2: more straightforward way of writing
func filterSimple(links []string, baseUrl string) []string {
	var ret []string
	for _, link := range links {
		if withSameBaseSimple(link, baseUrl) { // if withSameBase then only keep this link
			ret = append(ret, link)
		}
	}
	return ret
}

func withSameBaseSimple(link, baseUrl string) bool {
	return strings.HasPrefix(link, baseUrl)
}
