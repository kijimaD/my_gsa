package a

func f() {
	b()
}

func a() {} // want `a is unused`

func b() {

}
