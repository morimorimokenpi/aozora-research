package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Entity struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	InfoURL  string
	ZipURL   string
}

func findEntities(siteURL string) ([]Entity, error) {
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}

	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		println(elem.Text(), elem.AttrOr("href", ""))
	})

	return nil, nil
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entities, err := findEntities(listURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Println(entity.Title, entity.ZipURL)
	}
}
