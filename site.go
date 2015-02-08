package main

import (
	"fmt"
	"strings"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/env", env)
	
	fmt.Println("listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello World!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

func env(res http.ResponseWriter, req *http.Request) {
	env := os.Environ()

    fmt.Fprintln(res, "List of Environtment variables : \n")

    for index, value := range env {
       name := strings.Split(value, "=") // split by = sign

       fmt.Fprintf(res, "[%d] %s : %v\n", index, name[0], name[1])
    }
}
    


