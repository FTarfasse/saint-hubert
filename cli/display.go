package cli

import (
	c "../collect"
//	"fmt"
	"reflect"
)

func Columns(r c.Result) (sizes []int) {
	iter := reflect.ValueOf(r).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		sizes = append(sizes, v.Len())
	}
//	fmt.Println(sizes)
	return
}

// func Maxer(array [][]) []int {
// 	rand := []int{0, 1}
// 	return rand
// }