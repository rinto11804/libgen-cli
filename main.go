package main

import (
	"libgen/cmd"
	"libgen/pretty"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		color.Red("provide a book name")
		return
	}
	books, err := cmd.Search(&cmd.SearchOpt{
		Query: args[1],
	})
	if err != nil {
		log.Println(err)
		return
	}
	pretty.TablePrinter(books)
	selectedBook := pretty.Draw(books)
	if selectedBook != nil {
		err := cmd.Downloader(selectedBook)
		log.Println(err)
	}
}
