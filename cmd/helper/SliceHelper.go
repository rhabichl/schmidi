package helper

func RemoveValue(input []string, d string) []string {
	var result []string
	for _, v := range input {
		if v != d {
			result = append(result, v)
		}
	}
	return result
}

func GetFileByName(f []*Fi, name string) *Fi {
	for _, v := range f {
		if v.Name == name {
			return v
		}
	}
	return nil
}
