package accumulate

// Accumulate applies each entry of the string array to the function passed.
func Accumulate(text []string, fn func(string) string) []string {

	if fn == nil {
		return text
	}

	var result []string
	for _, t := range text {
		result = append(result, fn(t))
	}
	return result
}
