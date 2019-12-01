//package main
//
//import (
//	"io"
//	"os"
//)
//
//type Page struct {
//	title string
//	body  []byte
//}
//
//func (p *Page) Load(file string) {
//
//	b := make([]byte, 1024)
//
//	if file, fileOk := os.OpenFile(file, os.O_RDONLY, 0600); fileOk == nil {
//
//		for {
//
//			n, nOk := file.Read(b)
//
//			if nOk != nil || nOk == io.EOF {
//				break
//			}
//
//			b = append(b, b[:n]...)
//		}
//	}
//
//	p.body = b
//}
