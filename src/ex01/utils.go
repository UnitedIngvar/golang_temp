package main

import "Users/igorpetrov/GolandProjects/Go_Day01-0/src/ex01/models"

func CakeMapKeyExcept(
	baseMap map[string]models.Cake, toExcludeMap map[string]models.Cake) []string {
	var result []string
	for key := range baseMap {
		if _, contains := toExcludeMap[key]; !contains {
			result = append(result, key)
		}
	}

	return result
}

func IngredientMapKeyExcept(
	baseMap map[string]models.Ingredient, toExcludeMap map[string]models.Ingredient) []string {
	var result []string
	for key := range baseMap {
		if _, contains := toExcludeMap[key]; !contains {
			result = append(result, key)
		}
	}

	return result
}
