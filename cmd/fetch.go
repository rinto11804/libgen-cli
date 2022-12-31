package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

const (
	Host       = "libgen.is"
	SearchHref = "<a href='book.index.php.+</a>"
	SearchMD5  = "[a-zA-Z0-9]{32}"
	Fields     = "id,title,author,filesize,extension,md5,year,language,pages,publisher,edition,coverurl"
)

type Book struct {
	ID        string
	Title     string
	Author    string
	Publisher string
	Year      string
	Pages     string
	MD5       string
	Language  string
	Edition   string
	FileSize  string
	Extention string
	CoverUrl  string
}

type SearchOpt struct {
	Query        string
	SearchMirror url.URL
	Results      int
}

type Detailes struct {
	Hashes string
}

func Search(options *SearchOpt) ([]*Book, error) {
	options.Results = 25
	options.SearchMirror.Scheme = "https"
	options.SearchMirror.Host = Host
	options.SearchMirror.Path = "search.php"
	q := options.SearchMirror.Query()
	q.Set("req", options.Query)
	q.Set("lg_topic", "libgen")
	q.Set("open", "0")
	q.Set("view", "simple")
	q.Set("res", fmt.Sprint(options.Results))
	q.Set("phrase", "1")
	q.Set("column", "def")
	options.SearchMirror.RawQuery = q.Encode()

	res, err := getBody(options.SearchMirror.String())
	if err != nil {
		return nil, err
	}

	hashes := hashParser(res)
	books, err := getDetailes(hashes)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func hashParser(res []byte) string {
	var hashes string = ""
	re := regexp.MustCompile(SearchHref)
	matches := re.FindAllString(string(res), -1)

	for _, m := range matches {
		re := regexp.MustCompile(SearchMD5)
		hashes = hashes + re.FindString(m) + ","
	}
	return hashes
}

func getDetailes(hashes string) ([]*Book, error) {
	var books []*Book
	opt := &SearchOpt{
		SearchMirror: url.URL{
			Scheme: "https",
			Host:   Host,
			Path:   "json.php",
		},
	}
	q := opt.SearchMirror.Query()
	q.Set("ids", hashes)
	q.Set("fields", Fields)
	opt.SearchMirror.RawQuery = q.Encode()

	res, err := getBody(opt.SearchMirror.String())
	if err != nil {
		return nil, err
	}
	var responce []map[string]string
	json.Unmarshal(res, &responce)

	for i := range responce {
		book := &Book{
			ID:        responce[i]["id"],
			Title:     responce[i]["title"],
			Author:    responce[i]["author"],
			Publisher: responce[i]["publisher"],
			Year:      responce[i]["year"],
			Edition:   responce[i]["edition"],
			MD5:       responce[i]["md5"],
			Pages:     responce[i]["page"],
			Language:  responce[i]["language"],
			Extention: responce[i]["extension"],
			FileSize:  responce[i]["filesize"],
			CoverUrl:  Host + "/" + responce[i]["coverurl"],
		}
		books = append(books, book)

	}

	return books, nil
}
