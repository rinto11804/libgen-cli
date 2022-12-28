package cmd

import (
	"fmt"
	"regexp"
)

const (
	findKey = "http://62.182.86.140/main/.+(pdf|epub)"
	bookUrl = "http://library.lol/main/"
)

// http://62.182.86.140/main/2741000/e15f308bce0b4256721ef7e26c8a7062/R.%20Mark%20Volkmann%20-%20Svelte%20and%20Sapper%20in%20Action-Manning%20Publications%20Co..epub

func getDownloadURL(book *Book) (string, error) {
	res, err := getBody(bookUrl + book.MD5)
	if err != nil {
		return "error", err
	}
	downloadURL := regexp.MustCompile(findKey).FindString(string(res))
	return downloadURL, nil

}

func Downloader(book *Book) {
	dUrl, e := getDownloadURL(book)
	fmt.Println(dUrl, e)
}
