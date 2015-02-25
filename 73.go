package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher myFetcher) {
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

	for _, u := range urls {
		fmt.Printf("from: [%s], fetching: [%s]\n", url, u)
		Crawl(u, depth-1, fetcher)
	}

	return
}

func main() {
	Crawl("http://golang.org/", 1, fetcher)

	//time.Sleep(10000)

	for u, _ := range fetcher {
		//<-fetcher[u].done

		fmt.Println("fetcher[", u, "] = {\n   body\n   urls = [")
		for _, s := range fetcher[u].urls {
			fmt.Println("      ", s, ",")
		}
		fmt.Println("   ]\n}")
	}
}

type myFetcher map[string]*myResult

type myResult struct {
	body string
	urls []string
	done chan bool
}

func (f myFetcher) get_urls(url string) error {
	largeurls := make([]string, 1024)
	var scratch string = f[url].body
	var ret error = nil

	var i int = 0
	for {
		if i >= len(largeurls) {
			ret = fmt.Errorf("get_urls: exceded length of urls array at %d", i)
			break
		}
		n := strings.Index(scratch, "href=\"")
		if n >= 0 {
			scratch = scratch[n + 6:]
			en := strings.Index(scratch, "\"")
			if en >= 0 {

				if scratch[:en] == "" {
					scratch = scratch[en:]
					continue
				}

				sn := strings.Index(scratch[:en], "//")
				if sn == 0 {
					scratch = "http:" + scratch
					en = strings.Index(scratch, "\"")
				}

				hn := strings.Index(scratch[:en], "http:")
				if hn != 0 {
					hn = strings.Index(scratch[:en], "https:")
				}
				if hn != 0 {
					largeurls[i] = url + scratch[1:en]
					i++

				} else {
					largeurls[i] = scratch[:en]
					i++
				}
			}
			scratch = scratch[en:]

		} else {
			ret = nil
			break
		}
	}

	if (i > 0) {
		(*f[url]).urls = make([]string, i-1)
		(*f[url]).urls = largeurls[:i-1]
		(*f[url]).done = make(chan bool)

	} else {
		(*f[url]).urls = nil
	}

	//(*f[url]).done <- true

	return ret
}

func (f myFetcher) Fetch(url string) (string, []string, error) {

	_, ok := f[url]
	if ok {
		//return f[url].body, f[url].urls, nil
		return "", nil, fmt.Errorf("Fetch: %s: already fetched this url", url)
	}
	res, err := http.Get(url)
	if err != nil {
		return "", nil, fmt.Errorf("Fetch: %s: %s", url, err.Error())
	}
	var str_body string
	buff := make([]byte, 1024)
	for {
		n, err := res.Body.Read(buff)
	//fmt.Println("got res: ", res)
		if err != nil && err != io.EOF {
			return "", nil, fmt.Errorf("Fetch: %s: %s", url, err.Error())
		}
		if n == 0 {
			break
		}

		str_body += string(buff[:n])
	}
	res.Body.Close()

	// make sure the body is not empty before setting f[url]
	if str_body != "" {
		f[url] = &myResult{"", nil, nil}
		(*f[url]).body = str_body
		(*f[url]).done = make(chan bool)

	} else {
		return "", nil, fmt.Errorf("Fetch: %s: body was empty")
	}

	gu_err := f.get_urls(url)
	if gu_err != nil {
		fmt.Println(gu_err.Error())
	}
	return f[url].body, f[url].urls, nil
}

var fetcher myFetcher = make(myFetcher)

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
