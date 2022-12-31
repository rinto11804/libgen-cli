package main

import (
	"fmt"
	"libgen/cmd"
	"libgen/pretty"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("provide a book name")
		return
	}
	books, err := cmd.Search(&cmd.SearchOpt{
		Query: args[1],
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.TablePrinter(books)
	selectedBook := pretty.Draw(books)
	if selectedBook != nil {
		err := cmd.Downloader(selectedBook)
		if err != nil {
			fmt.Println(err)
		}
	}
}
