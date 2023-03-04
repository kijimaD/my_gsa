package a

import "fmt"

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

// this
// is
// safe
func safe() {
	// check multi line function comment
}

// ... // FIXME want `useless function comment!`
func warn() {
	// ↑assertしたいが、wantコメントの長さが含まれてしまうせいで検証できない
}
