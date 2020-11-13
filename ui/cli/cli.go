package cli

import (
	c "../../collect"
	"fmt"
	"strconv"
	"strings"
	"github.com/acarl005/stripansi"
)

const (
	Separator = "|"
	Space     = " "
	BackLine  = "\n"

	Reset = 0
	Bold  = 1
	Underline = 4
	ColorFormat = "\033[%v;%vm%s\033[0m" // reset / color number / message

	Black  = 30
	Red    = 31
	Green  = 32
	Yellow = 33
	Blue   = 34
	Purple = 35
	Cyan   = 36
	White  = 37

	BlackBackground  = 40
	RedBackground    = 41
	GreenBackground  = 42
	YellowBackground = 43
	BlueBackground   = 44
	PurpleBackground = 45
	CyanBackground   = 46
	WhiteBackground  = 47
)

var Headers = []string{
	"Link",
	"Status",
}

func ColumnsSizes(results []c.Result) (sizes []int) {

	if len(results) == 0 {
		panic("Can't calculate columns sizes on an empty array !")
	}

	maxAddress := len(results[0].Address)
	maxStatus := len(results[0].Status)
	maxCode := len(strconv.Itoa(results[0].Code))
	maxSource := len(results[0].Source)

	if len(results) > 1 {

		for i := 1; i < len(results); i++ {
			if maxAddress < len(results[i].Address) {
				maxAddress = len(results[i].Address)
			}
			if maxStatus < len(results[i].Status) {
				maxStatus = len(results[i].Status)
			}
			if maxCode < len(strconv.Itoa(results[i].Code)) {
				maxCode = len(strconv.Itoa(results[i].Code))
			}
			if maxSource < len(results[i].Source) {
				maxSource = len(results[i].Source)
			}
		}

	}

	sizes = append(sizes, maxAddress)
	sizes = append(sizes, maxStatus)
	sizes = append(sizes, maxCode)
	sizes = append(sizes, maxSource)

	return sizes
}

func DisplayCli(datas []c.Result) {
	maxs := ColumnsSizes(datas)
	colouredData := ColorOutput(datas)

	fmt.Printf("\033[1;45m%s\033[0m", fmt.Sprintf("\n SAINT-HUBERT FOUND THIS AT: %s \n\n", datas[0].Source))
	fmt.Print(BuildTable(colouredData, maxs[:2]))
}

func BuildTable(coloredData []c.Result, maxs []int) (formattedData []string) {
	headerBuilder := strings.Builder{}
	headerBuilder.WriteString(BuildHeader(Headers, Separator, maxs[:2]))
	headerBuilder.WriteString(BackLine)
	header := headerBuilder.String()
	formattedData = append(formattedData, header)

	for _, c := range coloredData { //
		lineBuilder := strings.Builder{}
		input := []string{
			c.Address,
			c.Status,
		}

		v := BuildLine(input, Separator, maxs[:2])
		lineBuilder.WriteString(v)

		lineBuilder.WriteString(BackLine)
		line := lineBuilder.String()

		formattedData = append(formattedData, line)
	}

	return formattedData
}

func BuildHeader(msg []string, separator string, size []int) (output string) {
	builder := strings.Builder{}
	builder.WriteString(separator)
	builder.WriteString(Space)
	builder.WriteString(msg[0])

	for i := 0; i < (size[0] - len(msg[0])); i++ {
		builder.WriteString(Space)
	}

	builder.WriteString(Space)
	builder.WriteString(separator)
	builder.WriteString(Space)
	builder.WriteString(msg[1])

	for j := 0; j < (size[1] - len(msg[1])); j++ {
		builder.WriteString(Space)
	}

	builder.WriteString(Space)
	builder.WriteString(separator)

	output = builder.String()

	return output
}

func BuildLine(msg []string, separator string, size []int) (output string) {

	builder := strings.Builder{}
	builder.WriteString(separator)
	builder.WriteString(Space)
	builder.WriteString(msg[0])

	for i := 0; i < (size[0] - len(msg[0])); i++ {
		builder.WriteString(Space)
	}

	builder.WriteString(Space)
	builder.WriteString(separator)
	builder.WriteString(Space)
	builder.WriteString(msg[1])
	escAnsiStr := stripansi.Strip(msg[1])

	for j := 0; j < (size[1] - len(escAnsiStr)); j++ {
		builder.WriteString(Space)
	}

	builder.WriteString(Space)
	builder.WriteString(separator)

	output = builder.String()

	return output
}

func ColorOutput(data []c.Result) []c.Result {

	for i := 0; i < len(data); i++ {
		if data[i].Code < 199 {
			data[i].Status = fmt.Sprintf(ColorFormat, Reset, White, data[i].Status)
		}
		if data[i].Code < 299 {
			data[i].Status = fmt.Sprintf(ColorFormat, Reset, Green, data[i].Status)
		}
		if data[i].Code < 399 {
			data[i].Status = fmt.Sprintf(ColorFormat, Reset, Yellow, data[i].Status)
		}
		if data[i].Code < 499 {
			data[i].Status = fmt.Sprintf(ColorFormat, Bold, Red, data[i].Status)
		}
		if data[i].Code < 599 {
			data[i].Status = fmt.Sprintf(ColorFormat, Reset, Cyan, data[i].Status)
		}
	}

	return data
}
