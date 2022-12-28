package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	host       = "libgen.is"
	searchHref = "<a href='book.index.php.+</a>"
	searchMD5  = "[a-zA-Z0-9]{32}"
	fields     = "id,title,author,filesize,extension,md5,year,language,pages,publisher,edition,coverurl"
)

type Book struct {
	ID          string
	Title       string
	Author      string
	Publisher   string
	Year        string
	Pages       string
	MD5         string
	Language    string
	Edition     string
	FileSize    string
	Extention   string
	CoverUrl    string
	DownloadUrl string
	Selected    bool
}

type SearchOpt struct {
	Query         string
	SearchMirror  url.URL
	Results       int
	Print         bool
	RequireAuthor bool
	Extention     []string
	Year          int
	Publisher     string
	Language      string
}

type Detailes struct {
	Hashes string
}

func Search(options *SearchOpt) ([]*Book, error) {
	var r int
	switch {
	case options.Results <= 25:
		r = 25
	case options.Results <= 50:
		r = 50
	default:
		r = 25
	}
	options.SearchMirror.Scheme = "https"
	options.SearchMirror.Host = host
	options.SearchMirror.Path = "search.php"
	q := options.SearchMirror.Query()
	q.Set("req", options.Query)
	q.Set("lg_topic", "libgen")
	q.Set("open", "0")
	q.Set("view", "simple")
	q.Set("res", fmt.Sprint(r))
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

func getBody(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 15,
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cant connect to %s error %s", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http returned witn status code %s", res.Status)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return b, nil
}

func hashParser(res []byte) string {
	var hashes string = ""
	re := regexp.MustCompile(searchHref)
	matches := re.FindAllString(string(res), -1)

	for _, m := range matches {
		re := regexp.MustCompile(searchMD5)
		hashes = hashes + re.FindString(m) + ","
	}
	return hashes
}

func getDetailes(hashes string) ([]*Book, error) {
	var books []*Book
	opt := &SearchOpt{
		SearchMirror: url.URL{Scheme: "https", Host: host, Path: "json.php"},
	}
	q := opt.SearchMirror.Query()
	q.Set("ids", hashes)
	q.Set("fields", fields)
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
			CoverUrl:  host + "/" + responce[i]["coverurl"],
		}
		books = append(books, book)

	}

	return books, nil
}
