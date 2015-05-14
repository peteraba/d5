package util

func SliceAppend(stringSlice []string, stringToAppend string) []string {
	result := []string{}

	for _, str := range stringSlice {
		result = append(result, str+stringToAppend)
	}

	return result
}
