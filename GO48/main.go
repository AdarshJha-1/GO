package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	// fmt.Printf("%+v\n", router)

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	router.HandleFunc("POST /hello/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("user_id")
		fmt.Println(userID)
		w.Write([]byte(fmt.Sprintf("Hello %s", userID)))
	})

	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err)
	}
}
