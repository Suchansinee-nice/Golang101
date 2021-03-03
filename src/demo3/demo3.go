package main

import "fmt"

var count int = 0

func main() {
	fmt.Println("Begin")

	//ประกาศตัวแปรแบบ explicit declaration คือ มีการระบุ type ชัดเจน ภาษา go เป็นภาษาที่ต้องระบุ type ชัดเจน
	// โดยจะเป็น static type คือไม่สามารถเปลี่ยนแปลง type ทีหลัง
	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "hello"
	var tmp3 bool = true

	// const tmp4 int = 0
	// tmp4 = 10

	//ประกาศตัวแปรแบบ implicit declaration การประกาศแบบ short form คือไม่ต้องพิมพ์ var หรือ data type คือให้มัน detect เองได้ว่าถ้า value แบบนี้ใช้ type อะไร
	//กรณีที่ต้องการประกาศตัวแปรแบบ implicit ต้องใส่เครื่องหมาย = เป็น :=
	tmp5 := 0
	tmp6 := "Hey"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp5)
	fmt.Println(tmp6)

	count++
	fmt.Println(count) //สามารถเปลี่ยนแปลงค่า global ได้
	run()
	run()
}

func run() {
	count++
	fmt.Println(count)
	// fmt.Println(tmp6) //undeclared

}
