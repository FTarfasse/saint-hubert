package main

import (
	c "./process/collect"
	cli "./output/cli"
	file "./output/file"
	"log"
	"os"
	exec "os/exec"
)

//const url = "http://example.com"
var ToFile bool

//const url = "http://perdu.com"
const url = "https://github.com/saint-hubert"

//--help
//--save
//--serve
//--all

func init() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	datas, err := c.Collect(url)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("\n \033[0;31m%v\033[0m\n", os.Args[1:])
	datas = c.Squeeze(datas, true)

	if ToFile {
		file.SaveToFile(datas)
	}

	//CLI
	cli.DisplayCli(datas)
}
