package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Chat")

	r := newRoom()

	http.Handle("/room", r)

	go r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
