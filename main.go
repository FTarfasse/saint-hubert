package main

import (
	get "./get"
	"fmt"
)

func main() {
	res := get.Retrieve("http://example.com")
	// fmt.Println(res.Address)
	// fmt.Println(res.Status)
	a := res.Address
	s := res.Status
	fmt.Printf("Link: %q | status: %q", a, s)
}
