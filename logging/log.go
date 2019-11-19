package main

import (
	"log"
	"time"
)

func init() {
	log.SetPrefix("LOGGING FROM CONSOLE: ")
	log.SetFlags(log.Lmicroseconds | log.Llongfile | log.Ldate)
	log.Println("starting logging:")
}

func main() {
	log.Println("1")
	log.Println("2")

	time.Sleep(time.Second)

	log.Println("3")
	log.Fatal("Exiting...")
}
