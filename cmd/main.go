package main

import "net/http"

func main() {
	// Initialize app

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
