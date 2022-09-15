package main

import(
	"fmt"
) 

func main() {

	var firstname string
	fmt.Println("enter your first name")
	fmt.scan(&firstname)

	var lastname string
	fmt.Println("enter your last name")
	fmt.scan(&lastname)


	fmt.Println("%v %v,firstname,lastname")
}