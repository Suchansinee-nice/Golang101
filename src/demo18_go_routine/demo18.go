package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ซื้อแว่น ที่ เซเว่น")
	fmt.Println("ซื้อนาฬิกา ที่ เซ็นทรัล")
	go fmt.Println("ซื้อผลไม้ ที่ สยามพารากอน")
	fmt.Println("ซื้อรถ ที่ ศูนย์ Toyota")
	time.Sleep(1 * time.Second)
}
