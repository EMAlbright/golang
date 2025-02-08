package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// describes how page data will be stored in memory
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600  file should be created with read-write
	// permissions for the current user only
	return os.WriteFile(filename, p.Body, 0600)
}

// load page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	/*
	 Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the
	 request path. This is because the path will invariably begin with "/view/", which is
	 not part of the page's title
	*/
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// writes to w, http.ResponseWriter
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
