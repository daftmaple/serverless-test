package handler

import (
	"fmt"
	"net/http"
)

// DocumentHandler ...
func DocumentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Document page</h1>")
}
