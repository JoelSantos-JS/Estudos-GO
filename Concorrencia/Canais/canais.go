package main

import (
	"fmt"
	"time"
	
)


func main() {
	canal := make(chan string)

	 go	escrever("Ola Mund" ,canal )
	 fmt.Println("Deposi da funcçao ")

	 mensagem :=  <-canal
	 fmt.Println(mensagem)

}


func escrever(texto string, canal chan string){
    time.Sleep(time.Second * 5)
	for i:= 0; i < 5; i++ {
		
		canal <- texto

		time.Sleep(time.Second)
	}
}