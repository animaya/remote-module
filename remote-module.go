package remotemodule

func Reverse(text string) string {
	result := ""
	for _, char := range text {
		result = string(char) + result
	}
	return result
}
