package a

import "fmt"

func f() {
	gopher := "a"                     // want `identifier is gopher!`
	aaa := "b"                        // OK
	fmt.Sprintf("%s %s", gopher, aaa) // want `identifier is gopher!`
}
