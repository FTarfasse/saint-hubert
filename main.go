package main

import (
	c "./collect"
	cli "./ui/cli"
	"log"
)

//const url = "http://example.com"
//
//const url = "http://perdu.com"
const url = "https://github.com/saint-hubert"

func main() {
	datas, err := c.Collect(url)
	if err != nil {
		log.Fatalln(err)
	}
	datas = c.Squeeze(datas)
	//fmt.Printf("Datas: %s", datas)

	//CLI
	cli.DisplayCli(datas)

	// JSON
	//j, err := json.Jsonify(datas)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Printf("Datas : %s", j)
}
