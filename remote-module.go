package remotemodule

func Reverse(text *string) {
	result := ""
	for _, char := range *text {
		result = string(char) + result
	}
}
