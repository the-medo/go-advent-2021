package utils

func SliceContians[T comparable](haystack []T, needle T) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
