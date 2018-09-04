package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

func linkScrape() {
    res, err := http.Get("https://en.wikipedia.org/wiki/Mineu_River")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)

    if err != nil {
        log.Fatal(err)
    }

    doc.Find("body a").Each(func(index int, item *goquery.Selection) {
        linkTag := item.Find("a")
        link, _ := linkTag.Attr("href")
        linkText := linkTag.Text()
        fmt.Printf("Link #%d: %s - %s\n", index, linkText, link)
    })
}

func main() {
    linkScrape()
}
