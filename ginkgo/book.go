package ginkgo

import "encoding/json"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func NewBookFromJSON(jsonStr string) (Book, error) {
	book := Book{}
	err := json.Unmarshal([]byte(jsonStr), &book)
	return book, err
}
