package main

import (
	"Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex01/models"
	"fmt"
	"strings"
)

func CompareRecipes(oldRecipes *models.Recipes, newRecipes *models.Recipes) error {
	oldCakesMap, err := getCakeMap(oldRecipes)
	if err != nil {
		return err
	}
	newCakesMap, err := getCakeMap(newRecipes)
	if err != nil {
		return err
	}

	comparisonResult, err := compareCakeMaps(oldCakesMap, newCakesMap)
	if err != nil {
		return err
	}

	fmt.Println(comparisonResult)
	return nil
}

func compareCakeMaps(
	oldCakeMap map[string]models.Cake,
	newCakeMap map[string]models.Cake) (string, error) {
	differencesBuilder := &strings.Builder{}

	deletedCakeNames := CakeMapKeyExcept(oldCakeMap, newCakeMap)
	for _, name := range deletedCakeNames {
		differencesBuilder.WriteString("REMOVED cake " + braces(name) + "\n")
	}

	for cakeName, newCake := range newCakeMap {
		if _, contains := oldCakeMap[cakeName]; !contains {
			differencesBuilder.WriteString("ADDED cake " + braces(cakeName) + "\n")
		} else {
			err := compareCakes(oldCakeMap[cakeName], newCake, differencesBuilder)
			if err != nil {
				return "", err
			}
		}
	}

	return differencesBuilder.String(), nil
}

func compareCakes(oldCake models.Cake, newCake models.Cake, differencesBuilder *strings.Builder) error {
	if oldCake.StoveTime != newCake.StoveTime {
		differencesBuilder.WriteString("CHANGED cooking time for cake " + braces(oldCake.Name) +
			" – " + braces(newCake.StoveTime) + " instead of " + braces(oldCake.StoveTime) + "\n")
	}

	oldIngredientMap, err := getIngredientsMap(oldCake)
	if err != nil {
		return err
	}
	newIngredientMap, err := getIngredientsMap(newCake)
	if err != nil {
		return err
	}

	compareIngredientMaps(oldIngredientMap, newIngredientMap, oldCake.Name, differencesBuilder)

	return nil
}

func compareIngredientMaps(
	oldIngredientsMap map[string]models.Ingredient,
	newIngredientsMap map[string]models.Ingredient,
	cakeName string,
	differencesBuilder *strings.Builder) {

	for ingredientName, newIngredient := range newIngredientsMap {
		if _, contains := oldIngredientsMap[ingredientName]; !contains {
			differencesBuilder.WriteString("ADDED ingredient " + braces(ingredientName) +
				" for cake " + braces(cakeName) + "\n")
		} else {
			compareIngredients(oldIngredientsMap[ingredientName], newIngredient, cakeName, differencesBuilder)
		}
	}

	deletedIngredientNames := IngredientMapKeyExcept(oldIngredientsMap, newIngredientsMap)
	for _, name := range deletedIngredientNames {
		differencesBuilder.WriteString("REMOVED ingredient " + braces(name) +
			" for cake " + braces(cakeName) + "\n")
	}
}

func compareIngredients(
	oldIngredient models.Ingredient,
	newIngredient models.Ingredient,
	cakeName string,
	differencesBuilder *strings.Builder) {

	if oldIngredient.Unit == "" && newIngredient.Unit != "" {
		differencesBuilder.WriteString(
			"ADDED unit " + braces(newIngredient.Unit) +
				" for ingredient " + braces(newIngredient.Name) +
				" for cake " + braces(cakeName) + "\n")
	} else if newIngredient.Unit == "" && oldIngredient.Unit != "" {
		differencesBuilder.WriteString(
			"REMOVED unit " + braces(oldIngredient.Unit) +
				" for ingredient " + braces(newIngredient.Name) +
				" for cake " + braces(cakeName) + "\n")
	} else if oldIngredient.Unit != newIngredient.Unit {
		differencesBuilder.WriteString(
			"CHANGED unit" +
				" for ingredient " + braces(newIngredient.Name) +
				" for cake " + braces(cakeName) + " – " +
				braces(newIngredient.Unit) + " instead of " + braces(oldIngredient.Unit) + "\n")
	}

	if newIngredient.Count != oldIngredient.Count {
		differencesBuilder.WriteString(
			"CHANGED unit count" +
				" for ingredient " + braces(newIngredient.Name) +
				" for cake " + braces(cakeName) + " – " +
				braces(newIngredient.Count) + " instead of " + braces(oldIngredient.Count) + "\n")
	}
}

func getCakeMap(recipes *models.Recipes) (map[string]models.Cake, error) {
	cakeMap := make(map[string]models.Cake, len(recipes.Cakes))

	for _, cake := range recipes.Cakes {
		if _, contains := cakeMap[cake.Name]; contains {
			return nil, fmt.Errorf("the provided database contains duplicate cake with name %v", cake.Name)
		}

		cakeMap[cake.Name] = cake
	}

	return cakeMap, nil
}

func getIngredientsMap(cake models.Cake) (map[string]models.Ingredient, error) {
	ingredientsMap := make(map[string]models.Ingredient, len(cake.Ingredients))

	for _, ingredient := range cake.Ingredients {
		if _, contains := ingredientsMap[ingredient.Name]; contains {
			return nil, fmt.Errorf("cake %v contains duplicate ingredient %v", cake.Name, ingredient.Name)
		}

		ingredientsMap[ingredient.Name] = ingredient
	}

	return ingredientsMap, nil
}

func braces(value string) string {
	return "\"" + value + "\""
}
