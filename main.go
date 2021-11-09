package main

import (
	"log"
	"main/render"
	"main/service"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = render.LoadFonts()
	if err != nil {
		log.Fatal(err)
	}

	lists := []render.List{
		{
			ID:    os.Getenv("TRELLO_MEALS_LIST"),
			Title: "Meal Plan",
		},
		{
			ID:    os.Getenv("TRELLO_TODO_LIST"),
			Title: "To Do",
		},
		{
			ID:    os.Getenv("TRELLO_SHOPPING_LIST"),
			Title: "Shopping",
		},
	}

	for i, l := range lists {
		cards, err := service.FetchCards(l.ID)
		if err != nil {
			log.Fatal(err)
		}
		lists[i].Cards = cards
	}

	events, err := service.GetCalendars(strings.Split(os.Getenv("ICS_URLS"), "|"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := render.Render(events, lists)
	ctx.SavePNG("./render.png")
}
