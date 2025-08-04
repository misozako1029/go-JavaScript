package main

import "fmt"

type Item struct{
	Name string
	Price float64
	Category  string
}

func taxRate(category string) float64{
	switch category{
	case "食品":
		return 0.08
	case "雑貨、書籍":
		return 0.10
	default:
		return 0.10
	}
}

func taxIncluded(price float64, rate float64) int{
	var prices float64=0
	prices=price*rate
	return int(float64(prices)+float64(price))
}

func printItem(item Item){
	fmt.Println("商品名:",item.Name)
	fmt.Println("カテゴリ:",item.Category)
	fmt.Printf("税率:%0.f%%\n",taxRate(item.Category)*100)
	taxrate:=taxRate(item.Category)
	fmt.Printf("税込価格:¥%d\n",taxIncluded(item.Price,taxrate))
	fmt.Println("----------------------------")
}

func main(){
	items:=[]Item{
		{Name:"パン",Price:100,Category:"食品"},
		{Name:"カレンダー",Price:200,Category:"雑貨、書籍"},
	}
	for _,i:=range items{
		printItem(i)
	}
}

