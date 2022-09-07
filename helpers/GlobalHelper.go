package helpers

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func LengthofArray(arr []string) int {
	length := len(arr)

	return length
}
