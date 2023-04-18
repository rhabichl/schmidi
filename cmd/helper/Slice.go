package helper

func RemoveValue(haystack []string, needle string) []string {
	var result []string
	for _, v := range haystack {
		if v != needle {
			result = append(result, v)
		}
	}
	return result
}
