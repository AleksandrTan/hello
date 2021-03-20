package main

import "fmt"

//return указатель на безимянную переменную
func newInt() *int {
	return new(int)
}

func getNewInc() {
	newIns := newInt() //объявление
	*newIns = 3500     // присвоение
	fmt.Println("Value - ", *newIns)
	fmt.Println("Adress - ", newIns)
}
