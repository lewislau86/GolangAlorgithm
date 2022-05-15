package main

import (
	"fmt"
	"reflect"
)
type X int 

func main() {
	var x X
	t:=reflect.TypeOf(x)
	fmt.Println(t.Name(),t.Kind(),t.String())
}
