package test

type ReturnTuple[T comparable, K comparable] struct {
	Val1 T
	Val2 K
}

type resetFunction = func()

func Stub[T any](target *T, stub T) resetFunction {
	oldTarget := *target
	*target = stub

	return func() {
		*target = oldTarget
	}
}
