package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello from app 2"))
	})
	http.ListenAndServe(":5000", nil)
}
