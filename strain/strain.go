package strain

// Ints defines a collection of int values
type Ints []int

// Lists defines a collection of arrays of ints
type Strings []string

// Strings defines a collection of strings
type Lists [][]int

// Keep filters a collection of ints to only contain the members where the provided function returns true.
func (i Ints) Keep(fn func(x int) bool) (o Ints) {
	for _, v := range i {
		if fn(v) {
			o = append(o, v)
		}
	}
	return
}

// Discard filters a collection to only contain the members where the provided function returns false.
func (i Ints) Discard(fn func(x int) bool) Ints {
	return i.Keep(func(x int) bool {
		return !fn(x)
	})
}

// Keep filters a collection of strings to only contain the members where the provided function returns true.
func (s Strings) Keep(fn func(x string) bool) Strings {
	l := make(Strings, 0)
	for _, v := range s {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

// Keep filters a collection of lists to only contain the members where the provided function returns true.
func (i Lists) Keep(fn func(x []int) bool) Lists {
	l := make(Lists, 0)
	for _, v := range i {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}
