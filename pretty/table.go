package pretty

import (
	"libgen/cmd"
	"time"

	"github.com/fatih/color"
)

const (
	tabs  = "   "
	lines = "-----------------------------------------------------------------------------------------"
)

var blue = color.New(color.FgBlue)
var m = color.New(color.FgMagenta).Add(color.Bold)

func TablePrinter(books []*cmd.Book) {
	for _, b := range books {
		blue.Println(lines)
		m.Print(b.ID)
		print(tabs, b.Title, "\n")
		println(tabs, b.Author, tabs, b.Year, tabs, b.Edition, tabs, b.Extention)
		println(b.MD5)
		blue.Println(lines)
		time.Sleep(time.Second * 1 / 3)
	}
}
func SelectorArrow(books []*cmd.Book) *cmd.Book {
	for _, b := range books {
		m.Print("* ", b.Title, "\n")
	}
	return nil
}
