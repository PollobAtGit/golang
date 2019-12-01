package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	title string
	body  []byte
}

func (p *Page) Load(file string) {

	if file, fileOk := os.OpenFile(file, os.O_RDONLY, 0600); fileOk == nil {

		for {

			b := make([]byte, 1024)

			n, nOk := file.Read(b)

			if nOk != nil || nOk == io.EOF {
				break
			}

			// fmt.Println(n, string(b[n:]))

			p.body = append(p.body, b[:n]...)
		}
	} else {
		fmt.Println(fileOk.Error())
	}
}

func loadPage(file string) (p Page) {
	p = Page{title: strings.TrimSuffix(file, filepath.Ext(file))}
	p.Load(file)
	return
}

func viewHandlerFunc(w http.ResponseWriter, r *http.Request) {
	p := loadPage(r.RequestURI[len(viewsUrlPath):])
	fmt.Fprint(w, string(p.body))
}

const viewsUrlPath = "/views/"

func main() {
	http.HandleFunc(viewsUrlPath, viewHandlerFunc)
	http.ListenAndServe(":8080", nil)

	// loadPage("welcome.html")
	// fmt.Println(string(loadPage("welcome.html").body))
}
