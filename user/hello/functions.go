package main

import "fmt"

//new()
//return указатель на безимянную переменную типа int
func newInt() *int {
	return new(int)
}

func getNewInc() {
	newIns := newInt() //объявление
	*newIns = 3500     // присвоение нового значения
	fmt.Println("Value - ", *newIns)
	fmt.Println("Address - ", newIns)
}

func celsRet() Cels {
	celsData := Cels(100.5)
	return celsData
}

func farRet() Far {
	farData := Far(100.5)
	return farData
}
