package main 

import "fmt"

func somar(n1 int, n2 int) int {

	return n1 + n2

}	

func main() {
	soma := somar(10,20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Joel deu certo")
}