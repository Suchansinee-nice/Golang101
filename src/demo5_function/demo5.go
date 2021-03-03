package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello world!")
	fn1()
	fn2("Good morning")
	fn3("Hi", 3)

	//single return
	const a, b int = 2, 3
	fmt.Printf("%d+%d =%d\n", a, b, sum(a, b)) //d = decimal
	// fmt.Println("a+b", sum(a, b)) //3
	// fmt.Println(sum(a, b)) //2
	// fmt.Println(sum(2, 3)) //1

	//multiple return
	var x, y int = swap(a, b) //x , y คือเก็บค่าที่รีเทินออกมาจาก function
	// x,y := swap(a,b) //short form
	fmt.Printf("%d , %d => %d,%d \n", a, b, x, y)

	x, y = swapV2(10, 20) //จะไม่ใช้ := เพราะเท่ากับเป็นการประกาศตัวแปร เราแค่ต้องการแทนค่าเข้าไปตัวแปรเดิมเฉยๆ เพราะมันประกาศไว้อยู้แล้วด้านบน คือ var x,y
	fmt.Printf("%d , %d => %d,%d", 10, 20, x, y)
}

func fn1() {
	fmt.Println("Hello world!")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

// func swapV2(a int, b int) ( x int, y int) {
// 	return b, a
// }

//short form
func swapV2(a, b int) (p, o int) { //naming ตัว p จะรีเทินใส่ x และ o รีเทินใส่ y
	o = a // ไม่ต้องใส่เป็น o := a เพราะ ไม่ได้ประกาศตัวแปรใหม่ คือการที่เราประกาศชื่อตอนรีเทิน ถือเป็นการสร้างตัวแปรมาอยู่แล้ว เราแค่เอาค่า a มา ใส่ใน o เป็นการ assign ค่าเข้าไปเฉยๆ ไม่ต้องใช้ :=
	p = b
	return
}
