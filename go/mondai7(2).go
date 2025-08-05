package main

import "fmt"

type Subject struct{
	Name string
	Score int
}

type Student struct{
	Name string
	Subjects []Subject
}

func (s Student)average() float64{
	total:=0
		for _,score:=range s.Subjects{
			total=score.Score+total
		}
	return float64(total)/float64(len(s.Subjects))
}


func main(){
	students:=[]Student{
    	{
			Name:"太郎",
			Subjects: []Subject{
				{Name:"国語",Score:80},
				{Name:"数学",Score:70},
				{Name:"英語",Score:90},//１つ目の構造体（subject構造体）を２つ目の構造体内に構造体のスライスとしておいて一つのフィールドのデータに複数の型を入力できるようになる
	
			},
		},
		{
			Name:"花子",
			Subjects: []Subject{
				{Name:"国語",Score:70},
				{Name:"数学",Score:85},
				{Name:"英語",Score:60},
			},
		},
	}
	for _,student:= range students{
		fmt.Println("名前:",student.Name)
		fmt.Printf("平均点:%.2f\n",student.average())
		if student.average()>=80.0{
			fmt.Println("判定:合格")
		}else{
			fmt.Println("判定:不合格")
		}
	}
	
}
