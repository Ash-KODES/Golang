package main

import "fmt"

type house struct {
	noRooms bool
	price   int
	city    string
}

func main(){

	var home=house{
		noRooms: false,
		price: 10,
		city: "bareilly",
	}
	fmt.Println(home)
}