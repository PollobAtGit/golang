package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type role int

const (
	_                    role = iota
	cook                 role = iota
	softwareEngineer     role = iota
	manager              role = iota
	chiefTechnicalOffice role = iota
)

// role is of type int/rune so any valid integer is assignable here
func (r role) strRep() string {

	weeks := []string{
		"Cook",
		"Software engineer",
		"Manager",
		"CTO",
	}

	i := int(r) - 1

	if i >= 0 && i < len(weeks) {
		return weeks[r-1]
	}

	return ""
}

func printRoles() {
	fmt.Println(cook.strRep(), softwareEngineer.strRep(), manager.strRep(), chiefTechnicalOffice.strRep())
}

func panicHandler() {
	if err := recover(); err != nil {
		log.Println("Panic occurred : ", err)
	}
}

func main() {

	for {
		fmt.Println("Choose your role:")

		fmt.Println("\t[1] ", cook.strRep())
		fmt.Println("\t[2] ", softwareEngineer.strRep())
		fmt.Println("\t[3] ", manager.strRep())
		fmt.Println("\t[4] ", chiefTechnicalOffice.strRep())

		fmt.Println("\tPress [-1] to quit")

		fmt.Println()

		stdIn := bufio.NewReader(os.Stdin)

		if content, err := stdIn.ReadString('\n'); err == nil {
			if v, convOk := strconv.Atoi(strings.TrimSuffix(content, "\r\n")); convOk == nil {

				if v == -1 {
					break
				}

				defer panicHandler()

				rStr := role(v).strRep()

				if rStr == "" {
					fmt.Println("You have chosen an invalid number")

				} else {
					fmt.Println("You have chosen :", rStr)
				}

				fmt.Println()
			}
		}
	}
}
