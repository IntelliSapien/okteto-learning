package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the hello world server 🚀")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World!")
}