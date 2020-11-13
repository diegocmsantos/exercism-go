package perfect

import "errors"

// Classification type
type Classification string

// ClassificationDeficient is deficient classification
const ClassificationDeficient Classification = "ClassificationDeficient"

// ClassificationPerfect is perfect classification
const ClassificationPerfect Classification = "ClassificationPerfect"

// ClassificationAbundant is abundant classification
const ClassificationAbundant Classification = "ClassificationAbundant"

// ErrOnlyPositive is an error for not positive number
var ErrOnlyPositive = errors.New("not a positive number")

func Classify(number int64) (Classification, error) {

  if number <= 0 {
    return "", ErrOnlyPositive
  }

	sum := int64(0)
	for i := int64(1); i <= number/2; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	if sum < number {
		return ClassificationDeficient, nil
	} else if sum > number {
		return ClassificationAbundant, nil
	}

	return ClassificationPerfect, nil
}
