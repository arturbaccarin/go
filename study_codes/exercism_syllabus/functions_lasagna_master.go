package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(sliceLayers []string, averageTime int) int {
	var numberLayers int

	if averageTime == 0 {
		averageTime = 2
	}

	numberLayers = len(sliceLayers) * averageTime

	return numberLayers
}

// TODO: define the 'Quantities()' function
func Quantities(sliceLayers []string) (int, float64) {
	var numberNoodles int = 0
	var sauceLiters float64 = 0

	for _, element := range sliceLayers {
		if element == "noodles" {
			numberNoodles += 50
		}

		if element == "sauce" {
			sauceLiters += 0.2
		}
	}

	return numberNoodles, sauceLiters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendSlice, yourSlice []string) {
	var lenFriendList int = len(friendSlice)
	var lenYourList int = len(yourSlice)

	var lastElementFriendList string = friendSlice[lenFriendList-1]

	yourSlice[lenYourList-1] = lastElementFriendList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountNeed []float64, numberPortions int) []float64 {
	var proportion float64 = float64(numberPortions) / 2.0
	var scaledQuantities []float64
	for i := 0; i < len(amountNeed); i++ {
		scaledQuantities = append(scaledQuantities, amountNeed[i]*proportion)
	}

	return scaledQuantities
}
