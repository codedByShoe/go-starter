package main

import "ABCD/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
