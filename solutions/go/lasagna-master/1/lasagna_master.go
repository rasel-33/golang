package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
        avgPreparationTime = 2
    }
    return len(layers) * avgPreparationTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    neededNoodles := 0
    neededSauce := 0.0
    for _, layer := range(layers) {
        if layer == "noodles" {
            neededNoodles += 50
        } else if layer == "sauce" {
            neededSauce += 0.2
        }
    }
    return neededNoodles, neededSauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, ownList []string) {
    ownList[len(ownList) - 1] = friendsList[len(friendsList) -1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64{
    calculatedQuantities := make([]float64, len(quantities))
	for i := 0; i < (len(quantities)); i++ {
        calculatedQuantities[i] = quantities[i] * float64(portion) / float64(2)
    }
    return calculatedQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
