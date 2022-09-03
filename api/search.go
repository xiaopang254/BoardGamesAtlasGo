package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

type BoardGameAtlas struct {
	//lower-case c in clientId means is private var
	clientId string
}

// Game
type Game struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	YearPublish uint   `json:year_published"`
	Description string `json:"description"`
	URL         string `json:"official_url"`
	ImageURL    string `json:"image_url"`
	RulesURL    string `json:"rules_url"`
}

type SearchResult struct {
	Games []Game `json:"games"`
	Count uint   `json:"count"`
}

// method in the BoardGameAtlas class
func (b BoardGameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) (*SearchResult, error) {

	//Create a HTTP client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)

	//check if there is an error
	if nil != err {
		// 	//returns error object
		return nil, fmt.Errorf("Cannot create http client: %v", err)
	}

	//Get query string object
	params := req.URL.Query()

	//populate the URL with query string params
	params.Add("name", query)
	params.Add("limit", fmt.Sprintf("%d", limit))
	params.Add("skip", strconv.Itoa(int(skip)))
	params.Add("client_id", b.clientId)

	//Encode query params and add it back to the reqest
	req.URL.RawQuery = params.Encode()

	fmt.Println("URL = %s\n", req.URL.String())

	//Make the call
	res, err := http.DefaultClient.Do(req)
	if nil != err {
		return nil, fmt.Errorf("Cannot create http client for invocation: %v", err)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Error Http status: %s", res.Status)
	}

	var result SearchResult
	//Deserialize the JSON payload to struct (for Go to read)
	if err := json.NewDecoder(res.Body).Decode(&result); nil != err {
		return nil, fmt.Errorf("Cannot serialize JSON payload: %v", err)
	}

	return &result, nil

}

func New(clientId string) BoardGameAtlas {
	//return clientId:clientId of boardGameatlas object
	return BoardGameAtlas{clientId: clientId}
}
