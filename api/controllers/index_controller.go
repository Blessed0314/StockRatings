package controllers

import(
	"net/http"
)

func GetData (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World 2"))
}