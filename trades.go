package main

import (
    "fmt"
    "strings"
    "net/http"
    "golang.org/x/net/html"
)


type NBA_Team struct {
    team string     // the team name
    abbrev string   // the city abbrev
    dates []string  // the list of dates of transcations
    transactions []string   // the list of transactions
}

// The url for getting city abbreviations for teams
const abbrevUrl = "https://en.wikipedia.org/wiki/Wikipedia:WikiProject_National_Basketball_Association/National_Basketball_Association_team_abbreviations"

// The url for getting the transaction data for teams
const espnUrl = "http://www.espn.com/nba/team/transactions/_/name/"

// A function type passed into baseScrape function in order to get certain data
type scrapeDest func(*html.Token, *html.Tokenizer, *NBA_Team)


// Aquire the transaction data for a nba team
func (ball *NBA_Team) get_data() {

    // We must first get the city abbrev in order to get the transaction data
    ball.get_abbrev()

    team_url := strings.Join([]string{espnUrl,ball.abbrev}, "")

    // We then scarpe ESPN for the transaction data
    ball.baseScrape(team_url, appendTransaction)
}

// In order to get the transaction data the city abbreviations for nba
//  teams is nedded. We get these abbreviations web scraping Wikipedia
func (ball *NBA_Team) get_abbrev(){

    // ESPN uses different city abbreviations for the Pelicans
    //  and the Jazz from the standard city abbreviations
    if ball.team == "New Orleans Pelicans" {
        ball.abbrev = "no"
    } else if ball.team == "Utah Jazz" {
        ball.abbrev = "utah"
    } else {
        ball.baseScrape(abbrevUrl, scrapeAbbreviation)
    }

}


// Scrape a url for td entries and use the scrapeDest function passed in
func (ball *NBA_Team) baseScrape(url string, fn scrapeDest) {

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error: Failed to crawl %s\n", url)
        return
    }

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

            // Check if the token is an 'td' tag
            isTableEntry := t.Data == "td"
            if isTableEntry {
                fn(&t, z, ball)
            }
        }
    }
}

// Get the transaction entry from ESPN and add the date and transaction
func appendTransaction(t *html.Token, z *html.Tokenizer, ball *NBA_Team) {
    if checkValidData(t) {

        // Date entry
        z.Next()
        date := strings.TrimSpace(html.UnescapeString(string(z.Text())))
        ball.dates = append(ball.dates, date)

        // closing td tag
        z.Next()
        //fmt.Println(z.Token())

        //beginning td tag
        z.Next()

        //text entry
        z.Next()
        trade := strings.TrimSpace(html.UnescapeString(string(z.Text())))
        ball.transactions = append(ball.transactions, trade)

        //closing td tag
        z.Next()
    }
}

// Get the abbreviation from Wikipedia webscraping
func scrapeAbbreviation(t *html.Token, z *html.Tokenizer, ball *NBA_Team) {
    var abbrev string

    // Check if the next token is the text token holding the abbreviation
    nextTag := z.Next()

    if nextTag == html.TextToken {
        abbrev = strings.TrimSpace(html.UnescapeString(string(z.Text())))

        // Find the following a tag whih holds the actual team name
        for{
            nextTag = z.Next()
            nextToken := z.Token()
            isAnchor := nextToken.Data == "a"
            if isAnchor {
                // Find the title attribute
                for _, a := range nextToken.Attr {
                    if a.Key == "title" {
                        // If the name refers to our nba team name
                        //  we have found the correct abbreviation
                        if a.Val == ball.team {
                            ball.abbrev = abbrev
                            return
                        }
                    }
                }
                break
            }

        }
    }
}

// For scraping ESPN we must check if the td tag is a transac entry
func checkValidData(t *html.Token) bool{
    ok := true
    for _, a := range t.Attr {
        if (a.Key == "colspan") || (a.Key == "width") {
            ok = false
        }
    }

    return ok

}
