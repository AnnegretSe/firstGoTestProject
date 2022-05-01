package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var a, c, b, sum int
	var response string

	a = 4
	b = 2
	c = 1
	sum = a * b * c

	response = "hello World " + strconv.Itoa(sum)

	//fmt.Println("Little go test")
	//fmt.Println(sum)

	http.HandleFunc("/helloworld", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, response)
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
