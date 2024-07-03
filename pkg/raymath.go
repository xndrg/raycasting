package raymath

func Sign(num float64) float64 {
	var result float64 = -1

	switch {
	case num > 0:
		result = 1
	case num == 0:
		result = 0
	}

	return result
}
