package a

import (
	"fmt"
)

type S struct {
	F1 string `label:"name"` // OK
	f1 string `label:"name"` // want `detect private field tag!`
}

func main() {
	s := S{
		"test",
		"test",
	} // OK
	fmt.Println(s)
}
