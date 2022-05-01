package main

import (
	"fmt"
	"html/template"
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

	//fmt.Println("Little go test") print a line on console
	//fmt.Println(sum)

	http.HandleFunc("/helloworld", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintln(writer, response) print a line fot http

		data := struct {
			TEXT string
		}{response}
		tmpl.Execute(writer, data) //execute a html template
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>{{.TEXT}}</h2>
</center></body></html>
`))
