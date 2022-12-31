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

var (
	blue = color.New(color.FgBlue)
	m    = color.New(color.FgMagenta).Add(color.Bold)
)

func TablePrinter(books []*cmd.Book) {
	println("\033[H\033[J")
	for _, b := range books {
		blue.Println(lines)
		m.Print(b.ID)
		print(tabs, b.Title, "\n")
		println(tabs, b.Author, tabs, b.Year, tabs, b.Edition, tabs, b.Extention)
		blue.Println(lines)
		time.Sleep(time.Second * 1 / 4)
	}
}
