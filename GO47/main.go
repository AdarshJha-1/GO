package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	/*
		h := HelloHandler{}
		if err := http.ListenAndServe(":8080", &h); err != nil {
			fmt.Println(err)
		}
	*/

	// or

	// if err := http.ListenAndServe(":8080", new(HelloHandler)); err != nil {
	// 	fmt.Println(err)
	// }

	// http.Handle("/hello", new(HelloHandler))	// instead of new(HelloHandler) we can use
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil { // will use default server mux
		fmt.Println(err)
	}
}
