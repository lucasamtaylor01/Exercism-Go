package chance

import (
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() (die_number int) {
	die_number = rand.Intn(20-1) + 1
	return
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() (wand_energy float64) {
	wand_energy = rand.Float64() * 12
	return
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := [8]string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	animalsSlice := animals[:]
	rand.Shuffle(len(animalsSlice), func(i, j int) {
		animalsSlice[i], animalsSlice[j] = animalsSlice[j], animalsSlice[i]
	})

	return animalsSlice
}
