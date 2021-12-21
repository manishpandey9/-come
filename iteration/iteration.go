package iteration

// In go there is only for loop there is no while, do , until keywords

func Repeat(character string, freq int) string {
	var repeated string
	for i := 0; i < freq; i++ {
		repeated += character
	}
	return repeated
}
