package test

type resetFunction = func()

func Stub[T any](target *T, stub T) resetFunction {
	oldTarget := *target
	*target = stub

	return func() {
		*target = oldTarget
	}
}
