package phptogo

func InArrayInt(array_val int, array []int) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}
func InArrayString(array_val string, array []string) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}
