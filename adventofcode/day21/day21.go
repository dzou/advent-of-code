package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Recipe struct {
	ingredients map[string]bool
	allergens   []string
}

func main() {
	file, _ := os.Open("adventofcode/day21/day21.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var recipes []Recipe
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " (contains ")

		ingredientsList := strings.Split(tokens[0], " ")
		ingredients := make(map[string]bool)
		for _, i := range ingredientsList {
			ingredients[i] = true
		}

		rightSide := strings.ReplaceAll(tokens[1], ")", "")
		allergens := strings.Split(rightSide, ", ")

		recipes = append(recipes, Recipe{ingredients, allergens})
	}

	fmt.Println(recipes)
	answer := solve1(recipes)
	fmt.Println(answer)
}

func solve1(recipes []Recipe) int {
	tracker := make(map[string]map[string]bool)

	for _, recipe := range recipes {
		for _, allergen := range recipe.allergens {
			allergenMap := tracker[allergen]
			if allergenMap == nil {
				tracker[allergen] = make(map[string]bool)
				for ingredient, _ := range recipe.ingredients {
					tracker[allergen][ingredient] = true
				}
			} else {
				for key, _ := range allergenMap {
					if !recipe.ingredients[key] {
						delete(allergenMap, key)
					}
				}
			}
		}
	}

	fmt.Println(tracker)

	//allIngredients := make(map[string]bool)
	//for _, recipe := range recipes {
	//  for i, _ := range recipe.ingredients {
	//    allIngredients[i] = true
	//  }
	//}

	usedIngredients := make(map[string]bool)
	for _, allergenMap := range tracker {
		for k, _ := range allergenMap {
			usedIngredients[k] = true
		}
	}

	count := 0

	for _, recipe := range recipes {
		for ingredient, _ := range recipe.ingredients {
			if !usedIngredients[ingredient] {
				count += 1
			}
		}
	}

	deduced := make(map[string]string)

	for len(tracker) > 0 {
		var singleton string

		for allergen, ingMap := range tracker {
			if len(ingMap) == 1 {
				for K, _ := range ingMap {
					deduced[allergen] = K
					singleton = K
				}

				delete(tracker, allergen)
				break
			}
		}

		for _, dict := range tracker {
			delete(dict, singleton)
		}
	}

	var allergens []string
	for k := range deduced {
		allergens = append(allergens, k)
	}

	sort.Strings(allergens)

	for _, s := range allergens {
		fmt.Print(deduced[s] + ",")
	}

	return count
}
