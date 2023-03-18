package a

import (
	_ "b" // import for b.Empty
	"context"
	"fmt"
)

var ctx context.Context

type I interface{ M() }

type T1 struct {
	ctx context.Context // want "struct include context!"
}

type T2 struct {
	fmt.Stringer
	// 埋め込みでインターフェースを満たしているため、スルー
	ctx context.Context // OK
}

type T3 struct {
	// ポインタのとき検知する
	ctx *context.Context // want "struct include context!"
}

type T4 struct {
	context.Context // OK
}

type T5 struct {
	I
	ctx context.Context // want "struct include context!"
}

type T6 struct {
	ctx context.Context // want "struct include context!"
}

func (*T6) M() {}

type T7 struct {
	// Stringer interface を満たすので、スルー
	ctx context.Context // OK
}

var _ fmt.Stringer = (*T7)(nil)

func (*T7) String() string { return "" }

type T8 struct {
	N       *int
	context *context.Context // want "struct include context!"
}

func (t *T8) M() {}

type T9 struct {
	// エクスポートしてるときはスルー
	CTX context.Context // OK
}
