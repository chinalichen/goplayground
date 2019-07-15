package main

import (
	"fmt"
	"time"
)

var h, o chan bool

func release() {
	for {
		for i := 0; i < 2; i++ {
			<-h
			fmt.Printf("H")
		}
		<-o
		fmt.Printf("O")
	}
}

func generateH2O(input string) {
	go release()
	for _, ch := range input {
		if string(ch) == "O" {
			o <- true
		} else if string(ch) == "H" {
			h <- true
		}
	}
}

func main() {
	h, o = make(chan bool, 150), make(chan bool, 150)
	generateH2O("HHHHHHOOO")
	time.Sleep(1 * time.Second)
}
