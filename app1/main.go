package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello from app 1"))
	})
	http.ListenAndServe(":3333", nil)
}
