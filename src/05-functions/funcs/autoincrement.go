package funcs

func Autoincrement() func() int {
	x := -1
	return func() int {
		x++
		return x
	}
}
