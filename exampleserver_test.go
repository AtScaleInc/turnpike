package turnpike_test

import (
	"net/http"

	"github.com/AtScaleInc/turnpike"
)

func ExampleServer_NewServer() {
	s := turnpike.NewServer(false)

	http.Handle("/ws", s.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
