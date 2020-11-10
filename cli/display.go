package cli

import (
	c "../collect"
	"reflect"
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

//func Columns(r c.Result) (sizes []int) {
func Columns(r c.Result) (c [][]int) {
	return c
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