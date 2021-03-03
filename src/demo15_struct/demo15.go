package main

import "fmt"

type customer struct { // การประกาศโครงสร้าง struct
	firstname string
	lastname  string
	code      int
	phone     string
}

func main() {
	cus := customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	} // การกำหนดค่าเริ่มต้น ให้ customer struct

	fmt.Println(cus)
	fmt.Println("name ", cus.firstname)
	fmt.Println("lastname", cus.lastname)

	cus.firstname = "Atikom"
	cus.lastname = "Sombutjalearn"

	fmt.Println("New name ", cus.firstname)
	fmt.Println("New lastname", cus.lastname)
}
