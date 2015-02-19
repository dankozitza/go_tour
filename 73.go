package main

import (
	"fmt"
	"net/http"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url)//body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

type myFetcher struct {}

func get_urls(body string) []string {
	var urls []string

	//fmt.Println(body)

	return urls
}

func (f myFetcher) Fetch(url string) (string, []string, error) {
	res, err := http.Get(url)
	b := make([]byte, res.ContentLength)
	res.Body.Read(b)
	//fmt.Println("got res: ", res)

	if (err == nil) {
		return string(b), get_urls(string(b)), nil
	}

	return "", nil, fmt.Errorf("fetching: %s returned: %s", url, err.Error())
}

var fetcher = myFetcher{}

// fetcher is a populated fakeFetcher.
//var fetcher = fakeFetcher{
//	"http://golang.org/": &fakeResult{
//		"The Go Programming Language",
//		[]string{
//			"http://golang.org/pkg/",
//			"http://golang.org/cmd/",
//		},
//	},
//	"http://golang.org/pkg/": &fakeResult{
//		"Packages",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/cmd/",
//			"http://golang.org/pkg/fmt/",
//			"http://golang.org/pkg/os/",
//		},
//	},
//	"http://golang.org/pkg/fmt/": &fakeResult{
//		"Package fmt",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//	"http://golang.org/pkg/os/": &fakeResult{
//		"Package os",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//}
