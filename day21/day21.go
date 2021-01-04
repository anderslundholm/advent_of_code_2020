package day21

import (
	"fmt"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func parseInput(input []string) ([]food, map[string]map[string]bool, map[string]map[string]bool, []string) {
	var foodList []food

	ingredientsAllergens := make(map[string]map[string]bool)
	allergensIngredients := make(map[string]map[string]bool)
	allergensIngredientsList := make(map[string][]string)
	var allIngredients []string
	for _, line := range input {

		foodItem := food{}
		splitLine := strings.Split(line, "(contains ")
		splitLine[1] = strings.TrimSuffix(splitLine[1], ")")
		foodItem.allergens = strings.Split(splitLine[1], ", ")
		splitLine[0] = strings.TrimSpace(splitLine[0])
		foodItem.ingredients = strings.Split(splitLine[0], " ")
		foodList = append(foodList, foodItem)
		for _, ingredient := range foodItem.ingredients {
			allIngredients = append(allIngredients, ingredient)
			ingredientsAllergens[ingredient] = make(map[string]bool)

			for _, allergen := range foodItem.allergens {
				allergensIngredientsList[allergen] = foodItem.ingredients
				allergensIngredients[allergen] = make(map[string]bool)
				ingredientsAllergens[ingredient][allergen] = true
				allergensIngredients[allergen][ingredient] = true
			}
		}

	}
	fmt.Println(ingredientsAllergens)
	return foodList, ingredientsAllergens, allergensIngredients, allIngredients
}

// func intersect(ingred2 map[string]bool, ingred1 []string) {
// 	for i := range ingred1 {
// 		if _, ok := ingred2[i]; !ok {
// 			delete(ingred1, i)
// 		}
// 	}
// }

func getAlergenFreeIngredients(input []string) int {
	allergenMatch := make(map[string]string)
	var result int

	_, ingredientsAllergens, allergensIngredients, _ := parseInput(input)

	for ingredient := range ingredientsAllergens {
		for _, allergen := range food.allergens {
			if ingredientWithAllergen, ok := allergensIngredients[allergen]; ok {
				intersect(ingredientWithAllergen, food.ingredients)
			}
		}
	}

	for _, allergens := range ingredientsAllergens {
		for allergen := range allergens {
			for allergen2, ingredients := range allergensIngredients {
				if allergen2 == allergen && len(ingredients) == 1 {
					// foodMap[allergen] = append(foodMap[allergen], )
					for k := range ingredients {
						allergenMatch[k] = allergen
						fmt.Println(k, allergen)
					}
				}
			}
		}
	}
	var count int
	for ingredient := range ingredientsAllergens {
		if _, ok := allergenMatch[ingredient]; !ok {
			fmt.Println(ingredient)
			// for _, ingredient1 := range allIngredients {
			// 	// for i := range allergensIngredients[allergen] {
			// 	if ingredient1 == ingredient {
			// 		fmt.Println(ingredient)
			// 		count++
			// 	}
			// 	// }
			// }

			// count++
		}
	}
	result = count

	return result
}
