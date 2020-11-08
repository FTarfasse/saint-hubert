package main

import (
	c "./collect"
	"fmt"
)

const url = "http://example.com"

func main() {
	//res := get.Retrieve("http://example.com")
	//// fmt.Println(res.Address)
	//// fmt.Println(res.Status)
	//a := res.Address
	//s := res.Status
	//fmt.Printf("Link: %q | status: %q", a, s)
	datas := c.Collect(url)

	for _, data := range datas {
		fmt.Printf("Link: %q | status: %q", data.Address, data.Status)
	}
}
