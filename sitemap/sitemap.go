package sitemap

import (
	"container/list"
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	link "github.com/Ishankhan21/go-site-map-generator/linkparser"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a diff domain
   5. Find all pages (BFS)
   6. print out XML
*/

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
}

func SiteMap() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	// maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")

	visitedURL := bfs(urlFlag, 10)
	encoder := xml.NewEncoder(os.Stdout)

	fmt.Println("Visited ++++", visitedURL)
	toXml := &urlset{}
	for _, l := range visitedURL {
		toXml.Urls = append(toXml.Urls, loc{Value: l})
	}

	encoder.Indent("  ", "    ")
	err := encoder.Encode(toXml)
	if err != nil {
		fmt.Println("Error Encoding XML:", err)
	}

	fmt.Println()
}

// Do Breath first search on the given URL and use the get() method below and return array of URLs
func bfs(urlPath *string, maxDepth int) []string {
	queue := list.New()
	queue.PushBack(urlPath) // Directly push the pointer to string

	visited := make(map[string]bool)
	// Initialize a slice to store the visited URLs
	var visitedUrls []string

	for queue.Len() > 0 {
		// Dequeue a URL
		element := queue.Front()
		queue.Remove(element)
		url, ok := element.Value.(*string)
		if !ok {
			fmt.Println("Error: Value is not a *string")
			continue
		}

		if visited[*url] {
			continue
		}
		visited[*url] = true
		visitedUrls = append(visitedUrls, *url)

		links := get(url)

		for _, link := range links {
			if maxDepth > 0 {
				queue.PushBack(&link)
			}
		}

		maxDepth--
	}

	return visitedUrls // Return the slice of visited URLs
}

func get(urlPath *string) []string {
	resp, err := http.Get(*urlPath)
	if err != nil {
		// return nil, err
	}
	defer resp.Body.Close()
	links, _ := link.Parse(resp.Body)
	fmt.Println("LInks ++++++", links)

	// Parse the URL from the response
	parsedURL, err := url.Parse(resp.Request.URL.String())
	if err != nil {
		// handle error
		fmt.Println("Error parsing URL:", err)
		// return
	}

	// Construct the base URL
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	fmt.Println("Base URL:", baseURL)

	var validUrls []string
	for _, l := range links {
		if strings.HasPrefix(l.Href, "/") {
			validUrls = append(validUrls, baseURL+l.Href)
		}
		if strings.HasPrefix(l.Href, "http") {
			validUrls = append(validUrls, l.Href)
		}
	}

	var currentDomainUrls []string
	for _, l := range validUrls {
		if strings.HasPrefix(l, baseURL) {
			currentDomainUrls = append(currentDomainUrls, l)
		}
	}

	return currentDomainUrls
}
