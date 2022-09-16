package main

import "fmt"

func main(){
	var number[10] int
	for i:=0;i<10;i++{
		number[i]=i+1
	}
	for i:=0;i<10;i++{
		fmt.Println(number[i])
	}
}