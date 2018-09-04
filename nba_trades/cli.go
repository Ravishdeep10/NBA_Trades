package nba_trades


import (
    "fmt"
    "log"
    //"bufio"
    //"os"
    //"strconv"

)

type CLI struct {
}



func (cli *CLI) Run(){
    //TODO use a web scraping method to get the teams name
    fmt.Println("What Team would you like to view?")

    teams := [30]string{
        "Boston Celtics", "Brooklyn Nets", "New York Knicks",
        "Philadelphia 76ers", "Toronto Raptors", "Chicago Bulls",
        "Cleveland Cavaliers", "Detroit Pistons", "Indiana Pacers",
        "Milwaukee Bucks", "Atlanta Hawks", "Charlotte Hornets", "Miami Heat",
        "Orlando Magic", "Washington Wizards", "Denver Nuggets",
        "Minnesota Timberwolves", "Oklahoma City Thunder",
        "Portland Trail Blazers", "Utah Jazz", "Golden State Warriors",
        "Los Angeles Clippers", "Los Angeles Lakers", "Phoenix Suns",
        "Sacramento Kings", "Dallas Mavericks", "Houston Rockets",
        "Memphis Grizzlies", "New Orleans Pelicans", "San Antonio Spurs"}

    for i := 0; i < 30; i++ {
        fmt.Println(i+1, ". ",teams[i])
    }


    team := NBA_Team{teams[getUserIn() - 1], []string{}, []string{}}
    fmt.Println(team)
    team.get_data()
}


func getUserIn() int {
    var num int
    fmt.Println("Enter team number: ")
    _, err := fmt.Scanf("%d", &num)
    if err != nil {
        log.Fatal(err)
    }
    return num
}
