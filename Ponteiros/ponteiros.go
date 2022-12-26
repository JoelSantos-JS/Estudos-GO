package main

import "fmt"


func main() {
	var var1 int = 10
	var var2 int = var1

	var1 += 2
	fmt.Println(var1,var2)

}