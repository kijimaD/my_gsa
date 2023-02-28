package a

import "fmt"

func f() {
	// ここでassertしたいが、assertコメントの長さに含まれてしまうのでうまく検証できない
	// aaa // can't assert...
	gopher := "a"                     // This is OK
	aaa := "b"                        // This is OK
	fmt.Sprintf("%s %s", gopher, aaa) // This is OK
}
