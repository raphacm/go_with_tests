package iteration

const repeatCount = 5

//Repeat repeats a string
func Repeat(letter string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}
	return repeated
}
