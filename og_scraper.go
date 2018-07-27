package main

import (
    "fmt"
    "os"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)

// Function to print to the user the usage of the program
func printUsage() {
    fmt.Println("Please put urls that you want to scarped in the command line")
    fmt.Println("Ex: go run og_scaper.go url1 url2 ... urln")
}

// Pull gref attribute from a token
func getHref(t html.Token) (href string, ok bool){

    for _, a := range t.Attr {
        if a.Key == "href" {
            href = a.Val
            ok = true
        }
    }

    return
}


//Extract all links from a given webpage
func webcrawl(url string, chUrls chan string, chFinished chan bool) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error: Failed to crawl %s\n", url)
        return
    }

    defer func() {
        chFinished <- true
    }()

    b := resp.Body
    defer b.Close()


    z := html.NewTokenizer(b)
    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End if Document
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check if the token is an 'a' tag
            isAnchor := t.Data == "a"
            if isAnchor {
                url, ok := getHref(t)
                if !ok {
                    continue
                }

                // Make sure the url has "http"
                hasProto := strings.Index(url, "http") == 0
                if hasProto {
                    chUrls <- url
                }
            }
        }
    }

}


func  main()  {
    if len(os.Args) == 1 {
        printUsage()
		os.Exit(1)
    }

    seedUrls := os.Args[1:]
    foundUrls := make(map[string]bool)

    //Channels
    chUrls := make(chan string)
    chFinished := make(chan bool)

    for _, url := range seedUrls{
        go webcrawl(url, chUrls, chFinished)
    }

    // Keep track of the channels
    for chCount := 0; chCount < len(seedUrls); {
        select {
        case url := <- chUrls:
            foundUrls[url] = true
        case <-chFinished:
            chCount++
        }
    }


    fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

	for url, _ := range foundUrls {
		fmt.Println(" - " + url)
	}

    close(chUrls)
}
