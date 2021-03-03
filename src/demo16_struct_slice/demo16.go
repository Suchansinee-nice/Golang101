package main

import "fmt"

type customer struct { // การประกาศโครงสร้าง struct
	firstname string
	lastname  string
	code      int
	phone     string
}

func main() {
	customerLists := []customer{}
	customer1 := customer{
		firstname: "Chaiyarin",
		lastname:  "Niamsuwan",
		code:      111990,
		phone:     "085661234",
	}
	customer2 := customer{
		firstname: "Atikom",
		lastname:  "Sombutjalearn",
		code:      111991,
		phone:     "085664321",
	}
	customer3 := customer{
		firstname: "Kritsana",
		lastname:  "Punyaphon",
		code:      111992,
		phone:     "085662344",
	}
	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	customerLists = append(customerLists, customer3)
	fmt.Println(customerLists)
}
