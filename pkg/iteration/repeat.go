package iteration

//Repeat repeats a string
func Repeat(letter string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += letter
	}
	return repeated
}
