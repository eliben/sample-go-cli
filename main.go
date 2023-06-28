package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello from sample-go-program")
	fmt.Println(quote.Go())

	username := "joemcj"
	fmt.Println("hash =", hash(username))
	fmt.Println("is authorized =", isAuthorized(username))
}
