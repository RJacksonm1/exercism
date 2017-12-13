// Package space handles immeasurable voids
package space

import "strings"

// Planet is an astronomical body orbiting a star or stellar remnant that is massive enough to be rounded by its own
// gravity, is not massive enough to cause thermonuclear fusion, and has cleared its neighbouring region of planetesimals.
type Planet string

// EarthYear in seconds
const EarthYear = 365.25 * 24 * 60 * 60

// Seconds in each planet's year
var planetYears = map[string]float64{
	"earth":   EarthYear,
	"mercury": EarthYear * 0.2408467,
	"venus":   EarthYear * 0.61519726,
	"mars":    EarthYear * 1.8808158,
	"jupiter": EarthYear * 11.862615,
	"saturn":  EarthYear * 29.447498,
	"uranus":  EarthYear * 84.016846,
	"neptune": EarthYear * 164.79132}

// Age calculates how old something is in a given planet's years
func Age(age float64, planet Planet) float64 {
	return age / planetYears[strings.ToLower(string(planet))]
}
