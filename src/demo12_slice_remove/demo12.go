package main

import (
	"fmt"
)

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6} //initial ค่า 1,2,3
	showSlice(numbers)

	//remove 2 แบบ  leading ด้านหน้า
	numbers = numbers[1:len(numbers)] //start,end ตำแหน่งเริ่มต้นที่ตรงไหน เริ้มเอาตั้งแต่ตัวที่1 คือ 2 ถึงจนสุดเลยก็คือขนาดสูงสุดของ slice คือ len(numbers) เขียนแบบนี้คือเอาที่เหลือ ไม่เอาตัวแรก
	showSlice(numbers)

	numbers = numbers[1:len(numbers)] //start,end
	showSlice(numbers)

	//tailing remmove ด้านหลัง
	numbers = numbers[0 : len(numbers)-1] // เริ่มต้นตั้งแต่ตัวแรก ใส่ 0 เอาจนถึง ทั้งหมดลบ1 คือ จะได้ผลลัพคือตัดตัวท้ายออกไป จะได้ที่เหลือคือ 0-5
	showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
