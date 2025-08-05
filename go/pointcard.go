package main

import "fmt"


func makeCounter(point int) func()int{
	pluspoint:=point
	return func()int{
		points:=pluspoint
		pluspoint++		
		return points
	}
}

func main(){
	m:=makeCounter(10)
	m2:=makeCounter(100)

	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m2())
	fmt.Println(m2())
}