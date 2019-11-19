package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readFile(file string) {

	reader, ok := os.Open(file)
	if ok == nil {

		for {
			b := make([]byte, 1024)
			n, e := reader.Read(b)

			if e != nil || n == 0 {
				break
			}

			fmt.Printf("=> %q", b[:n])
		}
	}
}

func readCsv(r io.Reader) {

	if r == nil {
		return
	}

	cr := csv.NewReader(r)

	for {

		line, e := cr.Read()

		switch e {
		case nil:
			fmt.Printf("%v\t%v", line[0], line[1])
		case io.EOF:
			return
		}
	}

}

func reader(f string) io.Reader {
	r, e := os.Open(f)

	if e == nil {
		return r
	}

	return nil
}

func main() {
	// readFile(`C:\Users\Ashiqul Islam Pollob\Downloads\data (1).json`)

	readCsv(reader("quiz.csv"))
}
