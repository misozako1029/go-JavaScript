package main 

import 	(
	"fmt"
	"errors"
)
func checkPrice(price int) error{

	if price <=0{
		return errors.New("価格が不正です")
	}
	return nil
}

func main(){
	c:=checkPrice(-100)
	if c!=nil{
		fmt.Println("エラー:",c)
	}else{
	fmt.Println("価格はOKです")
	}
}
