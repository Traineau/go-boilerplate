package controllers

import (
	"fmt"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
	return
}
