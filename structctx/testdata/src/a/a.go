package a

import (
	"context"
	"fmt"
)

type S struct {
	ctx context.Context // want `struct include context!`
	F1  string          // OK
}

func main() {
	s := S{context.Background(), "test"} // OK
	fmt.Println(s)
}
