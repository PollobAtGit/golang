package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func rootHTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func printURLHTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.RequestURI())
}

func handleViewHandler(w http.ResponseWriter, r *http.Request) {
	p := loadPage("wiki.html")
	fmt.Fprintf(w, string(p.body))
}

// Page ...
type page struct {
	title string
	body  []byte
}

func (p *page) save() error {
	return ioutil.WriteFile(p.title+".txt", p.body, 0600)
}

func (p *page) loadUsingIo(file string) {

	if f, ok := os.OpenFile(file, os.O_RDONLY, 0600); ok == nil {

		b := make([]byte, 1024)

		for {
			n, readOk := f.Read(b)

			if readOk != nil || readOk == io.EOF {
				break
			}

			p.body = append(p.body, b[:n]...)
		}
	}
}

func (p *page) load(file string) (ok error) {

	if content, ok := ioutil.ReadFile(file); ok == nil {
		p.body = content
	}

	return
}

func loadPage(file string) (p *page) {

	p = &page{title: strings.TrimSuffix(file, ".txt")}
	p.load(file)

	return
}

func main() {

	http.HandleFunc("/view", handleViewHandler)
	http.HandleFunc("/print-url", printURLHTTPHandler)
	http.HandleFunc("/", rootHTTPHandler)

	http.ListenAndServe(":8080", nil)
}
