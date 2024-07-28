package iteration

func Repeat(character string, uinput int) string {
	var result string
	for i := 0; i < uinput; i++ {
		result = character + result
	}
	return result
}
