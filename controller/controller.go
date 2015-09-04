package controller

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func StartServer() {
	http.HandleFunc("/", baseHandler)
	http.ListenAndServe(":8080", nil)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path[1:])

}
