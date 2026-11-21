package utils

func InArray(target string, strArr []string) bool {
	for _, element := range strArr {
		if target == element {
			return true
		}
	}
	return false
}

func IntInArray(target int64, intArr []int64) bool {
	for _, element := range intArr {
		if target == element {
			return true
		}
	}
	return false
}
