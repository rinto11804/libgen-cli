package cmd

import (
	"testing"
)

const downloadUrl string = "http://62.182.86.140/main/2741000/e15f308bce0b4256721ef7e26c8a7062/R.%20Mark%20Volkmann%20-%20Svelte%20and%20Sapper%20in%20Action-Manning%20Publications%20Co..epub"

var book = &Book{
	"2707413",
	"Svelte and Sapper in Action",
	"Mark Volkmann",
	"Manning Publications",
	"2020",
	"",
	"492CA877873F8EAC74F8A3A923C8FEEF",
	"English",
	"1",
	"25044281",
	"pdf",
	"libgen.is/2707000/492ca877873f8eac74f8a3a923c8feef-g.jpg",
}

func TestGetUrlDownloader(t *testing.T) {
	g, _ := getDownloadURL(book)
	t.Log(g)

}
