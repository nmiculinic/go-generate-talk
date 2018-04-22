package decorator

import (
	_ "context"
	xx "context"
)

//go:generate decorator-gen -in=$GOFILE -out=gen-$GOFILE

//@decorate:chan
func Increase(a int, ctx xx.Context) (int, int) {
	return a + 1, 0
}
