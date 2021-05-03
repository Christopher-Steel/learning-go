package iteration

func Repeat(s string, repetitions int) (result string) {
	for i := 0; i < repetitions; i++ {
		result += s
	}
	return
}