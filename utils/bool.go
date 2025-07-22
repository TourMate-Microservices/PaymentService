package utils

func IsBooleanRemain(input *bool, org bool) bool {
	// Remain
	if input == nil {
		return true
	}

	return *input == org
}
