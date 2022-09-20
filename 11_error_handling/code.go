package main

import(
	"fmt"
	"errors"
)


func main() {
	//err genrally used for error
	total,err:=sum(10,2)
	if err!=nil{
		fmt.Println("there is is error")
	}else{
		fmt.Println(total)
	}
}

func sum(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("start is greater than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total,nil
}
