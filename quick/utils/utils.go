package utils

func Optional[T interface{}](o T, fn ...func(T) T) T {
	for _, f := range fn {
		o = f(o)
	}
	return o
}
