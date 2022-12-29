package pretty

import (
	"fmt"
	"libgen/cmd"
	"strings"

	"github.com/manifoldco/promptui"
)

func Draw(books []*cmd.Book) *cmd.Book {
	searcher := func(input string, index int) bool {
		for _, book := range books {
			name := strings.Replace(strings.ToLower(book.Title), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			if strings.Contains(name, input) {
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

	_, title, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}
	for _, book := range books {
		if strings.Contains(title, book.Title) {
			return book
		}
	}

	return nil
}
