package main

import "net/http"
import "fmt"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("Hello World its running!"))
		fmt.Println("Received request")

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
