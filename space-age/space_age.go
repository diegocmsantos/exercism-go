package space

type Planet string

const earthSecondsInYear = 31557600

var earthBasedYear = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func (planet Planet) earthBasedYearInSeconds() float64 {
	return earthBasedYear[planet] * earthSecondsInYear
}

// Age calculate age for the given Planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / planet.earthBasedYearInSeconds()
}
