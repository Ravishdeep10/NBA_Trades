package nba_trades

import (
    "fmt"
    "strings"
    "unicode"
    "net/http"
    "golang.org/x/net/html"
)


type NBA_Team struct {
    team string
    dates []string
    transactions []string
}


func (ball *NBA_Team) get_abbrev() string{
    var b strings.Builder
	for _, char := range ball.team[:3] {
    		fmt.Fprintf(&b, "%c", byte(unicode.ToLower(char)))
	}
    return b.String()

}


func (ball *NBA_Team) get_data() {
    abbrev := ball.get_abbrev()
    ball.scrapeESPN(abbrev)
    fmt.Println(ball.dates)
    fmt.Println(ball.transactions)
}


func checkValidData(t html.Token) bool{
    ok := true
    for _, a := range t.Attr {
        if (a.Key == "colspan") || (a.Key == "width") {
            ok = false
        }
    }

    return ok

}


func (ball *NBA_Team) scrapeESPN(team string)  {
    espn_url := "http://www.espn.com/nba/team/transactions/_/name/"
    team_url := strings.Join([]string{espn_url,team}, "")

    fmt.Println(team_url)
    resp, err := http.Get(team_url)
    if err != nil {
        fmt.Println("Error: Failed to crawl %s\n", team_url)
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

                // Check if the token is an 'a' tag
                isTableEntry := t.Data == "td"
                if isTableEntry {
                    if checkValidData(t) {

                        // Date entry
                        z.Next()
                        //fmt.Println(z.Token())
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

                    // Make sure the url has "http"
                }
            }
        }

}
