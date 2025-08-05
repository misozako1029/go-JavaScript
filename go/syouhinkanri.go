package main

import "fmt"

type Item struct{
	Name string
	Price int
	Category string
}

func printAllItems(items []Item) {
	fmt.Println("【商品一覧】")
	for _,item:=range items{
		fmt.Printf("-%s (¥%d) [%s]\n",item.Name,item.Price,item.Category)
	}
}

func printItemsByCategory(items []Item,caregory string){
	fmt.Printf("【カテゴリ:%s の商品】\n",caregory)
	for _, item:=range items{
		if caregory==item.Category{
			fmt.Printf("-%s (¥%d)\n",item.Name,item.Price)
		}
	}
}





func main(){
	items:=[]Item{
		{"りんご",120,"食品"},
		{"Tシャツ",1500,"衣類"},
		{"ノート",200,"文房具"},
		{"パン",100,"食品"},
		{"ボールペン",150,"文房具"},
	}

printAllItems(items)
fmt.Println("----------------")
printItemsByCategory(items,"食品")
}
