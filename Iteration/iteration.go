package Iteration

func Repeat(character string) string {
	count := 5
	result := ""
	for i := 0; i < count; i++ {
		result += character
	}
	return result
}
