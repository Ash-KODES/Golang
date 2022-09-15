package main

import "fmt"

func main() {

	var firstname string
	fmt.Println("enter your first name")
	fmt.Scan(&firstname)

	var lastname string
	fmt.Println("enter your last name")
	fmt.Scan(&lastname)


	println(firstname+" "+lastname)

}