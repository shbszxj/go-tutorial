package tests

const repeatCount = 5

// Repeat returns 5 times of character
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
