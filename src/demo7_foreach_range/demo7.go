package main

import (
	"fmt"
)

func main() {
	//for loop ค่าที่มีอยู่ใน array
	courses := []string{"Andriod", "iOS", "React"} //ประกาศ array type string ใส่ค่าเข้าไปใน array

	//for ตามด้วยตัวแปร 2 ตัว ต้องมี 2 ตัว ใช้ชื่ออะไรก็ได้
	for index, item := range courses { //print course ข้างนออกมาภายใต้ for loop
		fmt.Printf("%d. %s \n", index+1, item)
	}

	//forr
	for _, item := range courses { // _ คือ ถ้าไม่ต้องการใช้ค่าบางอย่างก็คือใช้ _ ว่าไม่ได้ใช้ index ถ้าเราประกาศไปแต่ไม่ได้ใช้มันจะแดง เลยใส่ _
		fmt.Printf("%s \n", item)
	}

}
