package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time_per_layer int) (prep_time int) {
	if time_per_layer == 0 {
		prep_time = len(layers) * 2
	} else {
		prep_time = len(layers) * time_per_layer
	}
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 1
		}
		if layers[i] == "sauce" {
			sauce += 1
		}
	}

	noodles = noodles * 50
	sauce = sauce * 0.2

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, MyList []string) {
	MyList[len(MyList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, prop int) []float64 {
	newquantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		newquantities[i] = quantities[i] * float64(prop) / 2
	}
	return newquantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
