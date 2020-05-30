package space

import (
	"log"
)

type Planet string

/*
   Mercury: orbital period 0.2408467 Earth years
   Venus: orbital period 0.61519726 Earth years
   Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
   Mars: orbital period 1.8808158 Earth years
   Jupiter: orbital period 11.862615 Earth years
   Saturn: orbital period 29.447498 Earth years
   Uranus: orbital period 84.016846 Earth years
   Neptune: orbital period 164.79132 Earth years
*/

var numSecondsInEarthYear float64 = 31557600

var earthSecondsPerYear = map[Planet]float64{
	"Earth":   numSecondsInEarthYear,
	"Mercury": numSecondsInEarthYear * 0.2408467,
	"Venus":   numSecondsInEarthYear * 0.61519726,
	"Mars":    numSecondsInEarthYear * 1.8808158,
	"Jupiter": numSecondsInEarthYear * 11.862615,
	"Saturn":  numSecondsInEarthYear * 29.447498,
	"Uranus":  numSecondsInEarthYear * 84.016846,
	"Neptune": numSecondsInEarthYear * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	var secondsPerYear, ok = earthSecondsPerYear[planet]

	if !ok {
		log.Fatalf("%s is not a planet", planet)
	}

	return seconds / secondsPerYear
}
