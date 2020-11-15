// file package has the helper function for generating the output to a file
package file

import (
	c "../../collect"
	"fmt"
	errors "golang.org/x/xerrors"
	"time"
	"os"
	json "../json"
)

// SaveToFile transforms the content to JSON format and saves it to a file the function just created in the current
//directory of the program
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