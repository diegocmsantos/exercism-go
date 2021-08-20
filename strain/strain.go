package strain

type Ints []int
type Strings []string
type Lists [][]int

func (i Ints) Keep(fn func(x int) bool) Ints {
	if i == nil {
		return i
	}
	l := make(Ints, 0)
	for _, v := range i {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (i Ints) Discard(fn func(x int) bool) Ints {
	if i == nil {
		return i
	}
	l := make(Ints, 0)
	for _, v := range i {
		if !fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (s Strings) Keep(fn func(x string) bool) Strings {
	l := make(Strings, 0)
	for _, v := range s {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (s Strings) Discard(fn func(x string) bool) Strings {
	l := make(Strings, 0)
	for _, v := range s {
		if !fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (i Lists) Keep(fn func(x []int) bool) Lists {
	l := make(Lists, 0)
	for _, v := range i {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (i Lists) Discard(fn func(x []int) bool) Lists {
	l := make(Lists, 0)
	for _, v := range i {
		if !fn(v) {
			l = append(l, v)
		}
	}
	return l
}