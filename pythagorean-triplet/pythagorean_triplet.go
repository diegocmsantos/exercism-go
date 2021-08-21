package pythagorean

import "sort"

type Triplet [3]int

func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)
	c, m := 0, 2
	for c < max {
		for n := 1; n < m; n++ {

			a := m*m - n*n
			b := 2 * m * n
			c = m*m + n*n
			if (a < min || a > max) || (b < min || b > max) {
				continue
			}
			if c > max {
				break
			}

			arr := []int{a, b, c}
			sort.Ints(arr)
			result = append(result, Triplet{arr[0], arr[1], arr[2]})
		}
		m = m + 1
	}

	return result
}

func Sum(p int) []Triplet {
	result := make([]Triplet, 0)
	for i := 1; i <= p/3; i++ {
		for j := i + 1; j <= p/2; j++ {
			k := p - i - j
			if i*i+j*j == k*k {
				arr := []int{i, j, k}
				sort.Ints(arr)
				result = append(result, Triplet{arr[0], arr[1], arr[2]})

			}
		}
	}

	return result
}
