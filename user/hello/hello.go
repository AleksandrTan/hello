package main

import "fmt"

func main() {
	a, b := 1, 2
	var s string = "Hello"

	fmt.Println("Hello, world.", s)
	fmt.Printf("a = %d, b = %d \n", a, b)

	//Указатели
	getNewInc()

	fmt.Printf("Cels - %f\n", celsRet())
	fmt.Printf("Far - %f\n", farRet())

}
