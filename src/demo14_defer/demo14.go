package main

import (
	"fmt"
)

func main() {
	printFirst()
	// printFinish()
	defer printThird()  // -> ลำดับที่ 3
	defer printFourth() // -> ลำดับที่ 2
	defer printFinish() // -> ลำดับที่ 1
	printSecond()
}
func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}
func printThird() {
	fmt.Println("Third")
}
func printFourth() {
	fmt.Println("Fourth")
}
