package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/Ash-KODES"

func main() {

	fmt.Println("Web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("response's type : %T\n", response)

	defer response.Body.Close() //caller responibilty to close the response

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(data)

	fmt.Println(content)
}
