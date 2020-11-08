package main

import (
	c "./collect"
	d "./cli"
)

// const url = "http://example.com"
// const url = "http://perdu.com"
const url = "https://fr.wikipedia.org/wiki/Rebelles_(film,_2019)"

func main() {
	datas := c.Collect(url)
	d.DisplayCli(datas)
//	for _, data := range datas {
//		fmt.Printf("Link: %q | status: %q", data.Address, data.Status)
//	}
}
