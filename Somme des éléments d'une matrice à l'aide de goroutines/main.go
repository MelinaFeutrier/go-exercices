//Somme des éléments d'une matrice à l'aide de goroutines

package main

import (
	"fmt"
	"sync"
)

type Result struct {
	ID    int
	Somme int
}

func sum(slice []int, id int, wg *sync.WaitGroup, results chan Result) {
	defer wg.Done() 

	somme := 0
	for _, value := range slice {
		somme += value
	}

	results <- Result{ID: id, Somme: somme}
}

func main() {
	listes := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(listes))

	for id, slice := range listes {
		wg.Add(1)
		go sum(slice, id, &wg, results)
	}

	wg.Wait()
	close(results)

	total := 0
	for result := range results {
		fmt.Printf("Calcul %d : somme = %d\n", result.ID, result.Somme)
		total += result.Somme
	}

	fmt.Println("Toutes les sommes ont été calculées.")
	fmt.Printf("Le total de la matrice est : %d\n", total)
}
