package grains

import (
	"errors"
)

func Square(pos int) (uint64, error) {

	if pos < 1 || pos > 64 {
		return 0, errors.New("position must be between 1 and 64")
	}

	return 1 << uint(pos-1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
