package main

import (
	"fmt"
)

func main() {
	//ประกาศตัวแปรแบบ slice ด้วยการใช้ func make ระบุ อาเรย์ของ int
	var numbers1 = make([]int, 3, 5) //len =3 จะมีค่า default 0 มา 3 ตัว , cap = 5 , ถ้าใส่ len = 0 จะกลายเป็น array เปล่าๆที่ยังไม่มีค่า
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	// numbers1 = append(numbers1, 3) //cap=5 มันเกิน5 พอมันเกิน แทนที่จะเพิ่มทีละ1มันจะเพิ่มทีละ 5
	showSlice(numbers1) // 0 หมด เพราะ array มีการเพิ่มค่ามาแล้ว 3 ตัว จาก 5 ตัว โดยยัดค่า default = 0

	//ประกาศตัวแปร slice คือตัวแปรอารย์ที่ไม่มีการระบุ size (แบบไม่ใใช้ฟังชันmake), แบบไม่มีระบุความจุ สามารถทยอยเพิ่มได้ //จะไม่มีการจอง memory
	var numbers2 []int
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	// numbers2 = append(numbers2, 3)
	showSlice(numbers2)

}

//x[] int = slice คือตัวแปร อารย์ที่ไม่มีการระบุ size หรือ capacity
func showSlice(x []int) { //len = ฟังก์ชันโชว์ในอารเย์มีข้อมูลเท่าไร มันจุได้กี่ตัว cap = capacity คือความจุ ความสามารถจุได้เท่าไร จุสูงสุดได้เท่าไร
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) //%v ใช้ปริ้น value แบบ default ออกมาไม่ว่าจะ type ไหน int float string หรืออะไรก็ตาม กรณีที่เราไม่รู้ format ไหน ก็ใช้ %v
}
