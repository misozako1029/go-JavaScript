package main

import	( 
	"fmt"
	"errors"
)

type Item struct{
	Name string
	Price int
	Category string
}

func validateItem(item Item) error{
	if item.Name == ""{
		return errors.New("商品名が空です")
	}
	if item.Price <0 {
		return errors.New("価格が0以下です")
	}
	if item.Category == ""{
		return errors.New("カテゴリが指定されていません")
	}

	return nil
}

func main(){
		items:=[]Item{
		{"りんご",120,"食品"},
		{"Tシャツ",-100,"衣類"},
		{"",200,"文房具"},
		{"パン",100,"食品"},
		{"ボールペン",150,""},
	}
	for i,item:=range items{
		err:=validateItem(item)
		if err!=nil{
			fmt.Printf("商品%d:%s エラー:%s\n",i+1,item.Name,err)
		}else{
			fmt.Printf("商品%d:%sは正常です\n",i+1,item.Name)
		}

	}

}
