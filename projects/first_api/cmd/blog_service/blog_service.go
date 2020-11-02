package main

import (
	"log"
	"time"
)

func main() {
	log.Println("HelloWorld")
	for true {
		time.Sleep(5 * time.Millisecond)
	}
}
