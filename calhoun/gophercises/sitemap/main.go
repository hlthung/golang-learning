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

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
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
