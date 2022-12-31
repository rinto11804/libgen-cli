package pretty

import (
	"bufio"
	"fmt"
	"libgen/cmd"
	"os"
	"strings"
	"unicode"
)

func Draw(books []*cmd.Book) *cmd.Book {

	selectedIndex := 0

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("\033[H\033[J")
		fmt.Println("Enter j for down, k for up and ok for selection")
		for i, book := range books {
			if i == selectedIndex {
				fmt.Print("--> ")
			} else {
				fmt.Print("    ")
			}
			fmt.Println(book.Title)
		}

		input, _ := reader.ReadString('\n')
		input = strings.TrimRightFunc(input, unicode.IsSpace)

		if input == "k" {
			selectedIndex--
			if selectedIndex < 0 {
				selectedIndex = len(books) - 1
			}
		} else if input == "j" {
			selectedIndex++
			if selectedIndex >= len(books) {
				selectedIndex = 0
			}
		} else if input == "ok" {
			return books[selectedIndex]
		}
	}
}
