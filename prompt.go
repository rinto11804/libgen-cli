package main

import (
	"fmt"
	"libgen/fetch"
	"strings"

	"github.com/manifoldco/promptui"
)

func Draw(books []*fetch.Book) *fetch.Book {
	var selected *fetch.Book
	searcher := func(input string, index int) bool {
		for _, book := range books {
			name := strings.Replace(strings.ToLower(book.Title), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			if strings.Contains(name, input) {
				selected = book
				return true
			}
		}
		return false
	}
	var titles []string
	for _, book := range books {
		titles = append(titles, book.Title)
	}

	prompt := promptui.Select{
		Label:    "Title selector",
		Searcher: searcher,
		Items:    titles,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	fmt.Printf("You choose number %d: %s\n", i+1, books[i].Title)
	return selected
}
