package main

import (
	"fmt"
	"log"
	"main/service"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	lists := []string{
		"618858f063d5d166da25fb55",
		"614df2cdaa19637d9bc2062e",
		"615311beb3367261ab43acbe",
	}

	for _, l := range lists {
		cards, err := service.FetchCards(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cards)
	}
}
