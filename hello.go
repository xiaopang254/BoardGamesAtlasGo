package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/xiaopang254/BoardGamesAtlasGo/api"
)

// command to run go lang.
// command: go run .

func main() {
	// input format: bga --qeury "ticket to ride" --clientId abc123 --skip 10 --limit 5
	//flag.*datatype*(*name*, *default value*, *message*)
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame atlas clientid")
	limit := flag.Uint("limit", 10, "limit the number of results returned")
	skip := flag.Uint("skip", 0, "skip the number of results returned")
	timeout := flag.Uint("timeout", 10, "Timeout")

	//parse into command line
	flag.Parse()

	if isNull(*query) {
		log.Fatalln("Please use --query to set boardgame name to search")
	}

	if isNull(*clientId) {
		log.Fatalln("Please use --client to set boardgame atlas client_id")
	}

	//fmt.Printf("query=%s, clientid=%s, limit=%d, skip=%d  \n", *query, *clientId, *limit, *skip)

	//Create an instance of the BoardGame Atlas Client
	bga := api.New(*clientId)

	//Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel()

	//Make the invocation
	result, err := bga.Search(ctx, *query, *limit, *skip)
	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	//colors
	boldGreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()

	for _, g := range result.Games {
		fmt.Printf("%s \n", boldGreen("Name"), g.Name)
		fmt.Printf("%s: %s\n", boldGreen("Description"), g.Description)
		fmt.Printf("%s: %s\n\n", boldGreen("Url"), g.URL)
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
