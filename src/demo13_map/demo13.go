package main

import "fmt"

func main() {

	//ในวงเล็บ data type ของ key นอกวงเล็ก datat type value , initial เป็นรูปแบบ json
	//ไม่สามารถใส่ key ซ้ำกันได้
	var number1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("map : ", number1)
	fmt.Println("map 2 : ", number1["two"]) //ถ้าต้องการเฉพาะ value ใส่ key เข้าไป

	delete(number1, "one") //ลบ key , value
	fmt.Println("map : ", number1)

	//ประกาศ map แบบไม่ระบุค่า ไม่่ระบุขนาด ตอนแรก แต่ไปเพิ่มทีหลังได้
	var number2 = map[string]int{}
	number2["sky"] = 1
	number2["pink"] = 2
	fmt.Println("number2 : ", number2)

	//ประกาศ map แบบใช้ make ไม่ระบุค่า ไม่ระบุขนาด
	var number3 = make(map[string]int)
	fmt.Println("number3 : ", number3)

	//make map สร้างขึ้นมาโดยไม่กำหนดขนาดว่าต้องมีขนาดเริ่มต้นเท่าไร ใช้ make ได้
	//สร้าง color โดยที่ไม่กำหนดค่าเริ่มต้นให้มัน type string ทั้ง key,value
	//สามารถสร้างโดยไม่ต้องกำหนดชื่อ key,value มาก่อนตอนแรก ค่อยๆเพิ่มเข้าไปได้

	//ประกาศ map แบบใช้ make ไม่ระบุค่า ไม่ระบุขนาด แต่เพิ่มค่าทีหลัง
	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("colors : ", colors)
	fmt.Println("colors : ", colors["green"])

}
