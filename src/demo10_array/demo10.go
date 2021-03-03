package main

import (
	"fmt"
)

func main() {
	//[] เพื่อบอกว่าประกาศตัวแปร array ตามด้วย type
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{1, 2, 3, 4} //ไม่ใส่ type
	var array3 [3]string           //ยังไม่กำหนดค่า บอกก่อนว่าเป็น array ระบุ size เฉยๆ

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(len(array1))

	for _, item := range array1 {
		fmt.Print(item, ",")
	}

	for _, item := range array2 {
		fmt.Print(item, ",")
	}

	array3[0], array3[1], array3[2] = "android", "iOS", "React"
	for _, item := range array3 { // ถ้าวน loop แบบ index วน เอาต้องระวังเรื่อง scope ของ array ระวัง index out of bound
		fmt.Print(item, ",")
	}
}
