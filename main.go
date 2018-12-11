package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")
	var i interface{}
	i = "a"
	switch i.(type) {
	case int:
		fmt.Println(i)
	default:
		t := reflect.TypeOf(i)
		fmt.Println(t.String())
	}

}
