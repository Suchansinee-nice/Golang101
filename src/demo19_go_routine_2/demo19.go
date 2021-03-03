package main

import (
	"fmt"
	"time" // Import time มาเพื่อจับเวลาในการ Run
)

func buyGlassesAtSevenEleven() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
}
func buyWatchAtCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}
func buyFruitAtSiamParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}
func buyCarAtToyota() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ที่ศูนย์โตโยต้า : เสร็จแล้ว")
}
func main() {
	start := time.Now() // เริ่มจับเวลาในการ Run
	go buyGlassesAtSevenEleven()
	go buyWatchAtCentral()
	buyFruitAtSiamParagon()
	buyCarAtToyota()
	fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), " วินาที")
}
