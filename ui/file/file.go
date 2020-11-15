package file

import (
	c "../../collect"
	"fmt"
	errors "golang.org/x/xerrors"
	"time"
	"os"
	json "../json"
)

func SaveToFile(content []c.Result) {

	data, err := json.Jsonify(content)
	if err != nil {
		errors.Errorf("Couldn't generate Json for file: %s", err)
	}

	fileName := fmt.Sprintf("%s%v.json", "report_", time.Now().String())
	file, err := os.Create(fileName)
	if err != nil {
		errors.Errorf("Couldn't create file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		errors.Errorf("Couldn't write to file: %s", err)
	}

	fmt.Print("Json file created successfully !")
}