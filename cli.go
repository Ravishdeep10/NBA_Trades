package main


import (
    "fmt"
    "log"

)

// Basic Command Line Interface
type CLI struct {
}


// Give the user the prompt
func (cli *CLI) Run(){

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


    team := NBA_Team{teams[getUserIn() - 1], "", []string{}, []string{}}
    team.get_data()
    cli.Results(team)
}

// Output the results to the user
func (cli *CLI) Results(ball NBA_Team) {
    len := len(ball.dates)

	if len == 0 {
		fmt.Println("There is no recent transaction data for", ball.team)
		return
	}

	fmt.Println("The transactions for ", ball.team)
    for i := 0; i < len; i++ {
        fmt.Println(ball.dates[i], ":", ball.transactions[i])
    }
}
// Get the user's input
func getUserIn() int {
    var num int
    fmt.Println("Enter team number: ")
    _, err := fmt.Scanf("%d", &num)
    if err != nil {
        log.Fatal(err)
    }
    return num
}
