package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fnFor()
	fnWhile()
	fnWhileUsingBreak()
	fnContinue()
}

func fnFor() {
	for index := 0; index <= 10; index++ { //วนloop 0-10
		fmt.Printf("For Index %d\n", index) // \n เพื่อขึ้นบรรทัดใหม้ทุกครั้ง
	}
}

func fnWhile() { //มันไม่มีคำสั่ง while loop จะใช้คำสั่ง for แทน
	index := 0
	for index < 5 { //ใส่เงื่อนไข ทำเมื่อจริง จนกว่าจะเท็จแล้วออก
		index++                         //1-5
		fmt.Printf("Index %d\n", index) //ใส่ index++ ในวงลเ็บมันไม่ยอม ต้องไปใส่ข้างนอก
	}
}

//break

func fnWhileUsingBreak() { //มันไม่มีคำสั่ง while loop จะใช้คำสั่ง for แทน
	index := 0
	for true { //ใส่ true ให้วนทำตลอด
		if index > 5 {
			break
		}
		index++
		fmt.Printf("Break-Index %d\n", index) //ใส่ index++ ในวงลเ็บมันไม่ยอม ต้องไปใส่ข้างนอก
	}
}

//continue
func fnContinue() {
	sum := 0
	for i := 1; i < 7; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 12 (2+4+6)
}
