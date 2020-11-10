package main

import (
	c "./collect"
	d "./cli"
)

// const url = "http://example.com"
// const url = "http://perdu.com"
// const url = "https://fr.wikipedia.org/wiki/Rebelles_(film,_2019)"
const url = "https://github.com/saint-hubert"

func main() {
	datas := c.Collect(url)
	d.DisplayCli(datas)
	// infocolor := "\033[1;34m%s\033[0m"
	// fmt.Printf(infocolor, "HELLO WORLD")
}
