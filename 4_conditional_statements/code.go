package main

import "fmt"

func main(){

	var num1 int
	var num2 int
	
	fmt.Scan(&num1);
	fmt.Scan(&num2);
	if num1+num2 >= 10{

		fmt.Println(num1+num2)
	}else{

		fmt.Println("less that 10")
	}
}