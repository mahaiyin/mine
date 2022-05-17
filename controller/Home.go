package controller

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("mine's hello world..."))
}
