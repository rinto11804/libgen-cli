package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	FindKey = `http:\/\/62\.182\.86\.140\/main\/\d{7}\/\w{32}\/.+?(pdf|epub)`
	BookUrl = "http://library.lol/main/"
)

func newFile(f []byte, filename string) error {
	c, err := os.Create(filename)
	fmt.Println("Creating  file ", filename)
	if err != nil {
		return err
	}
	defer c.Close()
	fmt.Println("Writing file...... " + filename)
	c.Write(f)
	return nil
}

func getDownloadURL(book *Book) (string, error) {
	res, err := getBody(BookUrl + book.MD5)
	if err != nil {
		return "error", err
	}
	downloadURL := regexp.MustCompile(FindKey).FindString(string(res))
	return downloadURL, nil

}

func fetchBookFile(url string) ([]byte, error) {
	fmt.Println("Fetching file")
	res, err := getBody(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Downloader(book *Book) error {
	dURL, err := getDownloadURL(book)
	if err != nil {
		return err
	}
	if dURL == "error" {
		return errors.New("The url for book not gained")
	}
	f, err := fetchBookFile(dURL)
	if err != nil {
		return err
	}
	err = newFile(f, cleanFileName(book.Title+"."+book.Extention))
	return nil
}
func cleanFileName(filename string) string {
	temp := strings.Replace(filename, " ", "", -1)
	temp = strings.Replace(temp, "/", "_", -1)
	return "./" + strings.ReplaceAll(temp, ":", "_")
}
