package main

import "fmt"

type Subject struct{
	Name string     //教科名
	Score int		//点数
}

type Student struct{
	Name string		//生徒名
	Subjects []Subject //１つ目の構造体のデータを複数作る
}

func (s Student)average()float64{//平均点を出す関数
	total:=0
	for _,s:=range s.Subjects{
		total=s.Score+total
	}
	return float64(total)/float64(len(s.Subjects))
}

func rank(score int) string{//評価を出す関数
		switch {
		case score>=90:
			return "A"
		case score>=80:
			return "B"
		case score>=70:
			return "C"
		case score>=70:
			return "D"
		default:
			return "E"
		}

}

func comment(rank string)string{//評価に応じたコメントを出す関数
	switch rank{
		case "A":
			return"最高"
		case "B":
			return "よくできました"
		case "C":
			return "まあまあ"
		case "D":
			return"あと少し"
		case "E":
			return "頑張ろう"
		default:
			return "評価エラー"
	}
}

func main(){
	students:=[]Student{//データを追加
		{
			Name:"快司",
			Subjects:[]Subject{
				{Name:"国語",Score:70},
				{Name:"数学",Score:80},
				{Name:"英語",Score:80},
			},
		},
		{
			Name:"太郎",
			Subjects:[]Subject{
				{Name:"国語",Score:60},
				{Name:"数学",Score:90},
				{Name:"英語",Score:55},
			},
		},

	}
	for _,student:=range students{
	    fmt.Println("名前:",student.Name)//名前を表示
		for _,subjects:=range student.Subjects{//2つ目の構造体のSubjectsデータをそれぞれ一つずつ取り出す
			r:=rank(subjects.Score)//評価関数の呼び出し
			c:=comment(r)//コメント関数の呼び出し
			fmt.Printf("-%s:%d→評価:%s %s\n",subjects.Name,subjects.Score,r,c,)//文章を表示
		}
		fmt.Printf("平均点:%.2f\n",student.average(),)//平均点を表示
		fmt.Println("---------------------------------")//人ごとに区切る
	}
}