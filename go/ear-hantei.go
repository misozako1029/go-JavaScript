package main

import (
"fmt"
"errors"
)

func validateScore(score int) error{
	if score<0{
		return errors.New("スコアがマイナスです")
	}else if score>100{
		return errors.New("スコアが100を超えています")
	}
	return nil
}

func main(){
	err:=validateScore(150)	
	fmt.Printf("%s",err)
}
