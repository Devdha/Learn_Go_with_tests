package iteration

func Repeat(character string, repeat_count int) string {
	var result string
	for i := 0; i < repeat_count; i++ {
		result += character
	}
	return result
}
