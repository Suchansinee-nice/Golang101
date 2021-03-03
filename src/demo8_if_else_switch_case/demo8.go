package main

import "fmt"

func main() {
	fn1()
	switchCase()
	// grade()
}

func fn1() {
	value := 10
	if value == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!= 10")
	}

	if value > 10 || value < 2 {
		fmt.Println("value > 10 || value < 2")
	} else {
		fmt.Println("not value > 10 || value < 2")
	}

	if result := fn2(); result == "ok" {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK")
	}

}

func fn2() string {
	return "ok"
}

func switchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("something else")
		break
	}
}

func grade() {
	var score int
	fmt.Print("---GRADE--- \nEnter [Score]:")
	fmt.Scanf("%d", &score) //รับตัวเลขจำนวนเต็มที่กรอกเข้าเก็บไว้ในตัวแปร score
	fmt.Print("Your Grade =>")

	if score > 79 {
		fmt.Print("A")
	} else if score > 69 {
		fmt.Print("B")
	} else if score > 59 {
		fmt.Print("C")
	} else if score > 49 {
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
}
