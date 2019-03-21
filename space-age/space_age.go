// "Space Age" track
package space

type Planet string

// Earth orbital period in seconds
const EarthOrbitalPeriod = 31557600

// all planets and their orbital period
var planets = map[Planet]float64{
  "Earth":   EarthOrbitalPeriod,
  "Mercury": 0.2408467 * EarthOrbitalPeriod,
  "Venus":   0.61519726 * EarthOrbitalPeriod,
  "Mars":    1.8808158 * EarthOrbitalPeriod,
  "Jupiter": 11.862615 * EarthOrbitalPeriod,
  "Saturn":  29.447498 * EarthOrbitalPeriod,
  "Uranus":  84.016846 * EarthOrbitalPeriod,
  "Neptune": 164.79132 * EarthOrbitalPeriod,
}

// Age calculates how old someone would be on a given planet
// after a specified time in seconds
func Age(t float64, planet Planet) float64 {
  return t / planets[planet]
}
