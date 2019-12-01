package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Page ...
type Page struct {
	title string
	body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.title+".txt", p.body, 0600)
}

func (p *Page) loadUsingIo(file string) {

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

func (p *Page) load(file string) (ok error) {

	if content, ok := ioutil.ReadFile(file); ok == nil {
		p.body = content
	}

	return
}

func mainTest() {
	p := &Page{}

	p.loadUsingIo("modelpae.go")
	fmt.Println(string(p.body))

	page := &Page{}
	page.title = "new string"
	page.body = []byte("other screens")

	page.save()
}
