package main

import (
	"fmt"
	"time"
)


func main() {
	// CONCORRENCIA != PARALALELISMO

	go	Escrever("JOEL") // GOROUTINES
	Escrever("JOEL2")
}


func Escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}