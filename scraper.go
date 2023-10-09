package main 
 
import ( 
	// "encoding/csv" 
	"github.com/gocolly/colly" 
	// "log" 
	"flag"
	"strconv"
	"fmt"
	"time"
) 
 
type PlayerRecord struct { 
	rank, kills, matchesPlayed string
	playerName string
}


func singleThreadedQuerying() {
	var playerRecords []PlayerRecord
	c := colly.NewCollector()
	for i := 1; i <= 30; i++ {
		var link string ="https://battlefieldtracker.com/bf2042/leaderboards/stats/all/default?page=" + strconv.Itoa(i)
		c.Visit(link)
		c.OnHTML("tr", func(e *colly.HTMLElement) {
			callback(e, playerRecords, true)
		}) 
	}
}

func multiThreadedQuerying() {
	var playerRecords []PlayerRecord
	c := colly.NewCollector()
	for i := 1; i <= 30; i++ {
		var link string ="https://battlefieldtracker.com/bf2042/leaderboards/stats/all/default?page=" + strconv.Itoa(i)
		c.Visit(link)
		c.OnHTML("tr", func(e *colly.HTMLElement) {
			go callback(e, playerRecords, true)
		}) 
	}
}

func recordTime(fn func(), mode string) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed, "for", mode)
}

func callback(e *colly.HTMLElement, playerRecords []PlayerRecord, isGoRoutine bool) {
	playerRecord := PlayerRecord{}
	playerRecord.rank = e.ChildText("td.rank")
	playerRecord.playerName = e.ChildText("td.username")
	playerRecord.kills = e.ChildText("td.stat.highlight")
	playerRecord.matchesPlayed = e.ChildText("td.stat.collapse")
	playerRecords = append(playerRecords, playerRecord)
}

func main() {
	// sgf stands for "single-threaded querying first"
	stqfPtr := flag.Bool("stqf", false, "")
	flag.Parse() 
	if *stqfPtr == true {
		fmt.Println("Single-Threaded Querying First")
		recordTime(singleThreadedQuerying, "single-threaded querying")
		recordTime(multiThreadedQuerying, "multi-threaded querying")
	} else {
		fmt.Println("Multi-Threaded Querying First")
		recordTime(multiThreadedQuerying, "multi-threaded querying")
		recordTime(singleThreadedQuerying, "single-threaded querying")
	}
}