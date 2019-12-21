package php

// in_array
func InArrayInt(array_val int, array []int) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}

//in_array
func InArrayString(array_val string, array []string) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}

//array_count_values
func ArrayCountValues(array_list []string) (new_list map[string]int) {
	new_list = make(map[string]int)
	for _, v := range array_list {
		if _, ok := new_list[v]; ok {
			new_list[v]++
		} else {
			new_list[v] = 1
		}
	}
	return
}
