package funcs

func Mean(items ...float64) float64 {
	var result float64 = 0
	var iterations float64 = 0
	// with "_" index gets ignored
	for _, item := range items {
		result += item
		iterations += 1
	}

	return result / iterations

	// min09
}
