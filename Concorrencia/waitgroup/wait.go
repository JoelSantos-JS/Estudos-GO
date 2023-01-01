package main

import (
	"fmt"
	"time"
	"sync"
)


func main() {
	var wait  sync.WaitGroup

	wait.Add(2)

	go func() {
		escrever("hOLDS")
		wait.Done() /// -1
	 }()
		
	go func() {
		escrever("dssdsdsdsd")
		wait.Done() /// -1
	 }()
	
	 wait.Wait() // 0
}


func escrever(texto string){

	for i:= 0; i < 5; i++ {
		fmt.Println(texto)

		time.Sleep(time.Second)
	}
}