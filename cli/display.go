package cli

import (
	c "../collect"
	"strconv"
)

const (
	//PS1 = "\\[\\033[COLORm\\]"
	Message  = "Link: %s | status: %s "
	Prefix   = "\\033["
	DotPoint = ";"
	Sufix    = "m"

	Reset = 0
	// Bold = 1
	// Underline = 4

	Black = 30
	Red   = 31
	// Green  = 32
	GreenOutput = "\033[0;32m%s\033[0m"
	Yellow      = 33
	Blue        = 34
	Purple      = 35
	Cyan        = 36
	White       = 37

	BlackBackground  = 40
	RedBackground    = 41
	GreenBackground  = 42
	YellowBackground = 43
	BlueBackground   = 44
	PurpleBackground = 45
	CyanBackground   = 46
	WhiteBackground  = 47
)

func ColumnsSizes(results []c.Result) (sizes []int) {

	if len(results) == 0 {
		panic("Can't calculate columns sizes on an empty array !")
	}

	// init at results[0]
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

// func Maxer(array [][]) []int {
// 	rand := []int{0, 1}
// 	return rand
// }

//func DisplayCli(datas []c.Result) {
//	fmt.Printf("\n\tSAINT-HUBERT FOUND THIS AT: %s\n\n", datas[0].Source)
//	for _, data := range datas {
//		// fmt.Printf(formColor("Test", 200))
//		fmt.Printf(Message, data.Address, "OK")
//		fmt.Printf(" \033[1;32m%s\033[0m \n", data.Status)
//	}
//}
//
//func printWithColors(message string, link string, status string) {
//	fmt.Sprintf("Link: %s | status: %s", link, status)
//}
//
//func formColor(status string, color int) string {
//	// \033[
//	return Prefix + strconv.Itoa(Reset) + DotPoint + strconv.Itoa(color) + Sufix + status
//}

//
//func statusColor(){
//	if
//}
