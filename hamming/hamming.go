package hamming

import "errors"

// Distance will iterate over two strands and return the distance.
// distance means how many different letter two strands have for the same position.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("strands have different lenght")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
