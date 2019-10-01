package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiURL = "https://www.googleapis.com/books/v1/volumes"

var cli = http.Client{
	Timeout: time.Duration(5 * time.Second),
}

type Book struct {
	title string
}

type GoogleBook struct {
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type GoogleBookResponse struct {
	Items []GoogleBook `json:"items"`
}

type VolumeInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	PublishedOn string   `json:"publishedDate"`
	Description string   `json:"description"`
	Pages       int      `json:"pageCount"`
	Categories  []string `json:"categories"`
}

func main() {
	// get os args
	if len(os.Args) < 2 {
		log.Println("No book title given")
	}

	titles := make([]string, 0)

	for i := 1; i < len(os.Args); i++ {
		titles = append(titles, os.Args[i])
	}

	for _, title := range titles {
		searchtitle := strings.ReplaceAll(title, "-", "%20")
		queryString := fmt.Sprintf("?q=%s&maxResults=1", searchtitle)
		log.Printf("using queryString of %s", queryString)
		resp, err := cli.Get(apiURL + queryString)

		log.Printf("got status of %s", resp.Status)

		if err != nil {
			panic(err)
		}

		book := GoogleBookResponse{}

		err = json.NewDecoder(resp.Body).Decode(&book)

		if err != nil {
			panic(err)
		}

		if len(book.Items) != 1 {
			log.Println(book.Items)
			panic("bad length of items")
		}

		fil, err := os.Create(fmt.Sprintf("../books/%s.json", title))
		defer fil.Close()

		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(fil).Encode(&book.Items[0].VolumeInfo)

		if err != nil {
			panic(err)
		}
	}

}
