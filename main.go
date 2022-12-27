package main

import (
	"fmt"
	"libgen/fetch"
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
	opt := &fetch.SearchOpt{
		Query: args[1],
	}
	books, err := fetch.Search(opt)
	fmt.Println(books)
	if err != nil {
		log.Println(err)
		return
	}
	pretty.TablePrinter(books)
	selectedBook := Draw(books)
	Downloader(selectedBook)
}
