package iteration

func Repeat(character string, count int) string {
	var repeated_string string
	for i := 0; i < count; i++ {
		repeated_string += character
	}
	return repeated_string
}
