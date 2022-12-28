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
		color.Red("Please enter a book name")
		return
	}
	opt := &cmd.SearchOpt{
		Query: args[1],
	}
	books, err := cmd.Search(opt)
	if err != nil {
		log.Println(err)
		return
	}
	pretty.TablePrinter(books)
	selectedBook := pretty.Draw(books)
	if selectedBook != nil {
		cmd.Downloader(selectedBook)
	}

}
