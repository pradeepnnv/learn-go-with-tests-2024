package iteration

func Repeat(x string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += x
	}
	return repeated
}
