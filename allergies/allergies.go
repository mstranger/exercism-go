package allergies

// Allergens is the list of items
var Allergens = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies returns all allergens for a given score.
func Allergies(score uint) []string {
	result := make([]string, 0)

	for i := uint(0); i < uint(len(Allergens)); i++ {
		if score&(uint(1<<i)) != 0 {
			result = append(result, Allergens[uint(1<<i)])
		}
	}

	return result
}

// AllergicTo checks an item in the list of allergens.
func AllergicTo(score uint, item string) bool {
	for k, v := range Allergens {
		if item == v {
			return score&k != 0
		}
	}

	return false
}
