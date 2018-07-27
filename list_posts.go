package main

import (
    "fmt"
    "log"
    "net/http"

    //third party library
    "github.com/PuerkitoBio/goquery"
)

func postScrape() {
    // Request the HTML page.
    res, err := http.Get("http://jonathanmh.com")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)

    if err != nil {
        log.Fatal(err)
    }

    // use CSS selector found with browser inspector
    // for each, use index and item

    doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
        title := item.Text()
        linkTag := item.Find("a")
        link, _ := linkTag.Attr("href")
        fmt.Printf("Post #%d: %s - %s\n", index, title, link)
    })
}


func main() {
    postScrape()
}
