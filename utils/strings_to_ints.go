package utils

import "strconv"

func StringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))

	for i, s := range strs {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = val
	}

	return ints, nil
}
