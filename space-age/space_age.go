package space

type Planet string

const earthAgeInSeconds float64 = 31557600

func Age(seconds float64, planet Planet) float64 {

	planetMap := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	ageYears := seconds / earthAgeInSeconds

	return ageYears / planetMap[planet]
}
