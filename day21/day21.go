package day21

import (
	"sort"
	"strings"
)

// Food ...
type Food struct {
	Ingredients []string
	Allergens   []string
}

// Allergen ...
type Allergen struct {
	Name                string
	Ingredient          string
	PossibleTranslation map[string]struct{}
}

func parseInput(input []string) []*Food {
	var foodList []*Food
	for _, line := range input {
		foodItem := Food{}
		splitLine := strings.Split(line, "(contains ")
		splitLine[1] = strings.TrimSuffix(splitLine[1], ")")
		foodItem.Allergens = strings.Split(splitLine[1], ", ")
		splitLine[0] = strings.TrimSpace(splitLine[0])
		foodItem.Ingredients = strings.Split(splitLine[0], " ")
		foodList = append(foodList, &foodItem)
	}
	return foodList
}

func countAlergentFreeIngredients(ingredients map[string]int, allergens map[string]*Allergen) int {
	var result int
	for ingredient, sum := range ingredients {
		isAllergen := false
		for _, allergen := range allergens {
			if allergen.Ingredient == ingredient {
				isAllergen = true
				break
			}
		}
		if !isAllergen {
			result += sum
		}
	}
	return result
}

func (food *Food) intersect(allergens map[string]*Allergen, allergen string) map[string]struct{} {
	intersection := make(map[string]struct{})
	for _, ingredient1 := range food.Ingredients {
		for ingredient2 := range allergens[allergen].PossibleTranslation {
			if ingredient1 == ingredient2 {
				intersection[ingredient2] = struct{}{}
			}
		}
	}
	return intersection
}

func getAlergenFreeIngredients(input []string) int {
	allergens, ingredients := mapAllergens(input)
	translateAllergens(allergens)
	return countAlergentFreeIngredients(ingredients, allergens)
}

func mapAllergens(input []string) (map[string]*Allergen, map[string]int) {
	foodList := parseInput(input)
	allergens := make(map[string]*Allergen)
	ingredients := make(map[string]int)
	for _, food := range foodList {
		for _, ingredient := range food.Ingredients {
			ingredients[ingredient]++
		}
		for _, allergen := range food.Allergens {
			if _, ok := allergens[allergen]; !ok {
				allergens[allergen] = &Allergen{
					Name:                allergen,
					PossibleTranslation: make(map[string]struct{}),
				}
				for _, ingredient := range food.Ingredients {
					allergens[allergen].PossibleTranslation[ingredient] = struct{}{}
				}
				continue
			}
			allergens[allergen].PossibleTranslation = food.intersect(allergens, allergen)
		}
	}
	return allergens, ingredients
}

func translateAllergens(allergens map[string]*Allergen) {
	done := false
	for !done {
		done = true
		for _, allergen := range allergens {
			switch len(allergen.PossibleTranslation) {
			case 1:
				for ingredient := range allergen.PossibleTranslation {
					allergen.Ingredient = ingredient
				}
				for _, discoveredIngredient := range allergens {
					delete(discoveredIngredient.PossibleTranslation, allergen.Ingredient)
				}
			case 0:
				continue
			default:
				done = false
			}
		}
	}
}

func arrangeAllergens(input []string) string {
	allergens, _ := mapAllergens(input)
	translateAllergens(allergens)
	var allergenList []*Allergen
	for _, allergen := range allergens {
		allergenList = append(allergenList, allergen)
	}
	sort.Slice(allergenList, func(i, j int) bool {
		return allergenList[i].Name < allergenList[j].Name
	})
	var canonicalDangerousIngredientList []string
	for _, allergen := range allergenList {
		canonicalDangerousIngredientList = append(canonicalDangerousIngredientList, allergen.Ingredient)
	}
	return strings.Join(canonicalDangerousIngredientList, ",")
}
