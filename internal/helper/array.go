package helper

func ArrayFilter[T any](arr []T, callback func(T) bool) []T {
	var result []T
	for _, value := range arr {
		if callback(value) {
			result = append(result, value)
		}
	}
	return result
}

func ArrayFind[T any](arr []T, callback func(T) bool) (T, bool) {
	var toFind T
	for _, item := range arr {
		if callback(item) {
			return item, true
		}
	}
	return toFind, false
}

func ArrayFindIndex[T any](arr []T, callback func(T) bool) int {
	for i, item := range arr {
		if callback(item) {
			return i
		}
	}
	return -1
}
