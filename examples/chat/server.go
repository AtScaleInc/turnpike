package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AtScaleInc/turnpike"
)

func main() {
	s := turnpike.NewServer(false)
	http.Handle("/ws", s.Handler)
	http.Handle("/", http.FileServer(http.Dir("web")))

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
