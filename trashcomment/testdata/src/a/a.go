package a

import "fmt"

// ここでassertしたいが、assertコメントの長さに含まれてしまうのでうまく検証できない
// a // can't assert...

// This is OK
func f() {
	// aaa // This is OK
	// a
	gopher := "a"                     // This is OK
	aaa := "b"                        // This is OK
	fmt.Sprintf("%s %s", gopher, aaa) // This is OK

	dummy()
}

func dummy() {
	// not exist function comment
}
