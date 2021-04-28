package space

type Planet string

var multiplier = map[Planet]float64 {
    "Mercury": 0.2408467,
    "Venus": 0.61519726,
    "Earth": 1,
    "Mars": 1.8808158,
    "Jupiter": 11.862615,
    "Saturn": 29.447498,
    "Uranus": 84.016846,
    "Neptune": 164.79132,
}

// Returns age of person on a particular planet,
// given age in seconds and planet name
func Age(ageInSecond float64, planet Planet) float64 {
    earthAge := ageInSecond/float64(60*60*24*365.2425)
    return earthAge/multiplier[planet]
}
