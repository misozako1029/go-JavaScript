package main

import "fmt"

func makeMultiplier(x int) func (y int) int{
	return func (y int)int{
		return x*y
	}
}

func main(){
	multiplication:= makeMultiplier(6)
	multiplication2:=makeMultiplier(4)

	fmt.Println(multiplication(2))
	fmt.Println(multiplication2(8))
}