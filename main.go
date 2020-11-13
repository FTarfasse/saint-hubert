package main

import (
	c "./collect"
	errors "golang.org/x/xerrors"
	"os"
	"time"
	//cli "./ui/cli"
	json "./ui/json"
	"fmt"
	//exec "os/exec"
	//"os"
	"log"
)

const url = "http://example.com"
//const url = "http://perdu.com"
//const url = "https://github.com/saint-hubert"

func init() {
	//c := exec.Command("clear")
	//c.Stdout = os.Stdout
	//c.Run()
}

func main() {
	datas, err := c.Collect(url)
	if err != nil {
		log.Fatalln(err)
	}

	datas = c.Squeeze(datas, false)

	//CLI
	//cli.DisplayCli(datas)

	// JSON
	j, err := json.Jsonify(datas)
	if err != nil {
		log.Fatalln(err)
	}
	// FILE HANDLING
	fileName := fmt.Sprintf("%s_%v.json", "json_", time.Now().String())
	file, err := os.Create(fileName)
	if err != nil {
		errors.Errorf("Couldn't create file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString(j)
	if err != nil {
		errors.Errorf("Couldn't write to file: %s", err)
	}

	fmt.Printf("Datas : %s", j)
}
