package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/encoding/japanese"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors(author_id, TEXT, author TEXT, PRIMARY KEY(author_id))`,
		`CREATE TABLE IF NOT EXISTS contents(author_id, TEXT, title_id TEXT, title TEXT, content TEXT, PRIMARY KEY(author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	b, err := os.ReadFile("ababababa.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)

	res, err := db.Exec(`INSERT INTO contents(author_id, title_id, title, content) values (?, ?, ?, ?)`,
		"000879",
		"14",
		"あばばばば",
		content,
	)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := res.LastInsertId()
}
