package service

import (
	"net/http"
	"sort"

	"github.com/apognu/gocal"
)

func GetCalendars(urls []string) ([]gocal.Event, error) {
	events := make([]gocal.Event, 0)
	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		c := gocal.NewParser(response.Body)

		err = c.Parse()
		if err != nil {
			return nil, err
		}

		events = append(events, c.Events...)
	}

	sort.Slice(events, func(p, q int) bool {
		return events[p].Start.Before(*events[q].Start)
	})

	return events, nil
}
