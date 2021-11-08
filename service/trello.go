package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

type Card struct {
	Name string    `json:"name"`
	Due  time.Time `json:"due"`
}

func FetchCards(listId string) ([]Card, error) {
	var params url.Values = map[string][]string{
		"key":   {os.Getenv("TRELLO_KEY")},
		"token": {os.Getenv("TRELLO_TOKEN")},
	}
	u := fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?%s", listId, params.Encode())

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cards []Card
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	sort.Slice(cards, func(p, q int) bool {
		return cards[p].Due.Before(cards[q].Due)
	})

	return cards, nil
}
