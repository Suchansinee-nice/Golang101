package main

import (
	"fmt"
)

func main() {
	msg := "message"

	//1. declare pointer varaible
	//ประกาศตัวแปร pointer type string , assign ค่าให้ตัวแปร pointer มันจะ complie error
	//เพราะตัวแปร pointer ไม่สามารถเก็บค่าทั่วๆไปได้ เช่น string int ใดๆได้
	//เพราะมันไม่ได้มีหน้าที่เก็บค่า ตัวแปร pointer คือตัวชี้ ใช้ในการเก็บตำแหน่งของตัวแปร หรือตำแหน่งของ address ในการบรรจุค่าๆนึงอยู่
	//pointer ใช้ในกรณี performance tuning เพื่อที่ได้ชี้ไปที่ตำแหน่งของตัวแปรเหล่านั้นเพื่อเปลี่ยนแปลง value

	// var msgPointer *string = msg //error
	var msgPointer *string = &msg //ให้แสดง address ของตัวแปรต้องใส่ & นำหน้า

	fmt.Println(msg)

	//2.จะพิมพ์มาเป็นเลขฐาน 16 คือ adddress หรือตำแหน่งของตัวแปร pointer ตัวนี้
	//3. บางครั้งเราต้องการแค่รู้ตำแหน่งก่อน ไม่ได้ต้องการค่ามันมา แต่ต้องการรู้ที่อยู่ของค่านั้น เพื่อเราจะได้ไปเปลี่ยนค่า ญ ตำแหน่งนี้จากที่ไหนก็ได้โดยการ pass pointer หรือส่ง pointer ไปยังจุดต่างๆ
	fmt.Println(msgPointer)  //แสดงตำแหน่ง pointer
	fmt.Println(*msgPointer) //4.ใส่เครื่องหมาย * เพื่อดึง value

	changeMessage(&msg) //5.แทนที่รีเทินตัวแปรกลับไป แต่จะทำการเปลี่ยนแปลงค่าตัวแปร msg แทน, ส่งตำแหน่ง pointer หรือ ตัวแปรที่เก็บค่าตำแหน่ง pointer
	fmt.Println(msg)

	changeMessage2(&msg, "new mesage1") //&msg ตำแหน่ง pointer
	fmt.Println(msg)

	changeMessage2(msgPointer, "new mesage2") //ตัวแปร msgPointer ที่ประกาศไว้ ส่ง ตำแหน่ง pointer เหมือนกัน
	fmt.Println(msg)                          //ตัวแปร msg ที่ถูกเปลี่ยนแปลงค่าแล้ว

	changeMessage2(msgPointer, "new mesage3")
	fmt.Println(*msgPointer) //เทียบเท่ากับอ้างถึงค่าที่อยุ่ในตัวแปร pointer ใส่เครื่องหมาย * เพื่อดึง value

}

func changeMessage(aPointer *string) { //รับ address มา ใส่ตัวแปร aPointer แต่จะเอา aPointer = "" ไม่ได้ ต้อง *aPointer เพื่ออ้างถึงค่าข้างใน เปลี่ยนแปลงค่าข้างในที่ตัวแปร pointerชี้อยู่
	*aPointer = "New message" //ต้องการเปลี่ยนค่าที่ตัวแปร pointer ชี้อยู่, ทำให้ตัวแปร msg เปลี่ยนจาก message เป็น new message โดยที่เราไม่ต้องรีเทินค่ากลับไป
}

func changeMessage2(aPointer *string, newMessage string) { //รับ address มา
	*aPointer = newMessage //ต้องการเปลี่ยนค่าที่ตัวแปร pointer ชี้อยู่
}
