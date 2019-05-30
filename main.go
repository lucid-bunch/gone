package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/manifoldco/promptui"
)

type app struct {
	Client http.Client
}

type response struct {
	Result result `json:"result"`
}

type result struct {
	Document []document `json:"document"`
}

type document struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func newApp() *app {
	return &app{
		Client: http.Client{},
	}
}

const (
	baseURI = "http://api.ne.se/search?q"
	t       = "encyklopedi.lång"
	s       = "5"
)

func main() {
	app := newApp()
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid string")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Query",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	app.search(app.Client, result)
}

func (*app) search(client http.Client, query string) {
	req, reqErr := http.NewRequest("GET", baseURI+url.QueryEscape(query)+"&src="+url.QueryEscape(t)+"&size="+s, nil)
	if reqErr != nil {
		// do something
	}
	req.Header.Set("Accept", "application/json")
	resp, respErr := client.Do(req)
	if respErr != nil {
		// do something
	}

	response := &response{}
	json.NewDecoder(resp.Body).Decode(response)
	fmt.Printf("\n")
	for i, d := range response.Result.Document {
		if i < 9 {
			fmt.Print(" ")
		}
		fmt.Printf("%d. %s\n    ", i+1, d.Title)
		for j := 0; j < len(d.Title); j++ {
			fmt.Print("*")
		}
		fmt.Printf("\n    %s\n\n", d.Summary)
	}
	return
}
