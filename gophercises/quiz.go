package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader, e := os.Open("problems.csv")

	if e == nil {

		csvReader := csv.NewReader(reader)
		totalCorrectAnswerCount := 0

		for {

			// can't work with '"1"+1?'
			// TODO: configure

			columns, e := csvReader.Read()

			if e != nil || e == io.EOF {
				break
			} else {
				fmt.Printf("%v\t", columns[:1][0])

				readString, e := bufio.NewReader(os.Stdin).ReadString('\n')

				if e == nil {
					answer, ae := strconv.Atoi(strings.TrimSuffix(readString, "\r\n"))
					rightAnswer, re := strconv.Atoi(columns[1:][0])

					if ae == nil && re == nil {
						if answer == rightAnswer {
							totalCorrectAnswerCount++
						}
					}
				}
			}
		}

		fmt.Printf("Total correct answer count: %v", totalCorrectAnswerCount)
	}
}
