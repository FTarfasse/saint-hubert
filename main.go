package main

import (
	c "./collect"
	cli "./ui/cli"
	exec "os/exec"
	"os"
	"log"
)

const url = "http://example.com"
//const url = "http://perdu.com"
//const url = "https://github.com/saint-hubert"

func init(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	datas, err := c.Collect(url)
	if err != nil {
		log.Fatalln(err)
	}

	datas = c.Squeeze(datas, false)

	//CLI

	cli.DisplayCli(datas)

	// JSON
	//j, err := json.Jsonify(datas)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Printf("Datas : %s", j)
}