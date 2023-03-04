package a

import (
	"fmt"
)

func f() {
	many1("a", "b", "c", "d", "e")
	many2("a", "b", "c", "d", "e")
	safe1("a", "b", "c", "d")
	safe2("a", "b", "c", "d")
	no()
}

func many1(a, b, c, d, e string) { // want `too many arguments!`
	fmt.Print(a, b, c, d, e)
}

func many2( // want `too many arguments!`
	a,
	b,
	c,
	d,
	e string,
) {
	fmt.Print(a, b, c, d, e)
}

func safe1(a, b, c, d string) { // OK
	fmt.Print(a, b, c, d)
}

func safe2( // OK
	a,
	b,
	c,
	d string,
) {
	fmt.Print(a, b, c, d)
}

func no() {} // OK
