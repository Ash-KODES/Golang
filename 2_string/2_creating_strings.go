package main

import "fmt"

func main() {

	//explicit data type deination
	var name string
	name = "Ashutosh"

	//implicit data type defination
	var college = "IET"

	fmt.Println(name)

	fmt.Println(college)
}

//there can be only one main function in a package--https://forum.golangbridge.org/t/main-redeclared-in-this-block-warning/16231
