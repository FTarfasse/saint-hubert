package cli

import (
	c "../collect"
	"fmt"
	"reflect"
)

const (
	//PS1 = "\\[\\033[COLORm\\]"
	Message = "Link: %s | status: %s %s"
	Prefix  = "\033["
	Sufix   = "m"

	Reset = 0

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

func Columns(r c.Result) (sizes []int) {
	iter := reflect.ValueOf(r).MapRange()

	for iter.Next() {
		// k := iter.Key()
		v := iter.Value()
		sizes = append(sizes, v.Len())
	}

	return
}

// func Maxer(array [][]) []int {
// 	rand := []int{0, 1}
// 	return rand
// }

func DisplayCli(datas []c.Result) {
	fmt.Printf("\n\tSAINT-HUBERT FOUND THIS AT: %s\n\n", datas[0].Source)
	for _, data := range datas {
		fmt.Printf(Message, data.Address, data.Status, "\n")
	}
}

// func printWithColors(message string, link string, status string) {
// 	fmt.Sprintf("Link: %s | status: %s", link, status)
// }

//
//func formColor(color string) string{
//	return Prefix + color + Sufix
//}
//
//func statusColor(){
//	if
//}
