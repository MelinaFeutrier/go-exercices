package main

import "fmt"

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 3, "c": 4, "d": 5}

	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)

	result := fusionMaps(map1, map2)

	fmt.Println("Fusion:", result)
}

func fusionMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key, value := range map2 {
		
		result[key] += value
	}

	return result
}
