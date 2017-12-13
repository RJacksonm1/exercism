// Package space handles immeasurable voids
package space

import (
	"strings"
)

// Planet is an astronomical body orbiting a star or stellar remnant that is massive enough to be rounded by its own
// gravity, is not massive enough to cause thermonuclear fusion, and has cleared its neighbouring region of planetesimals.
type Planet string

// Seconds in each planet's year
const (
	EarthYear   = 31557600
	MercuryYear = EarthYear * 0.2408467
	VenusYear   = EarthYear * 0.61519726
	MarsYear    = EarthYear * 1.8808158
	JupiterYear = EarthYear * 11.862615
	SaturnYear  = EarthYear * 29.447498
	UranusYear  = EarthYear * 84.016846
	NeptuneYear = EarthYear * 164.79132
)

// Age calculates how old something is in a given planet's years
func Age(age float64, planet Planet) float64 {

	switch strings.ToLower(string(planet)) {
	case "mercury":
		age /= MercuryYear

	case "venus":
		age /= VenusYear

	case "earth":
		age /= EarthYear

	case "mars":
		age /= MarsYear

	case "jupiter":
		age /= JupiterYear

	case "saturn":
		age /= SaturnYear

	case "uranus":
		age /= UranusYear

	case "neptune":
		age /= NeptuneYear
	}

	return age
}
