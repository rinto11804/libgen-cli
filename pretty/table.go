package pretty

import (
	"libgen/fetch"
	"time"

	"github.com/fatih/color"
)

const (
	tabs  = "   "
	lines = "-----------------------------------------------------------------------------------------"
)

var blue = color.New(color.FgBlue)
var m = color.New(color.FgMagenta).Add(color.Bold)

func TablePrinter(books []*fetch.Book) {
	for _, b := range books {
		blue.Println(lines)
		m.Print(b.ID)
		print(tabs, b.Title, "\n")
		println(tabs, tabs, "@author", b.Author, tabs, "@year", b.Year, tabs, "@edition", b.Edition, tabs, "@", b.Extention)
		blue.Println(lines)
		time.Sleep(time.Second * 1 / 2)

	}

}
