//Recherche parallèle d’un élément dans un tableau avec Goroutines

package main

import (
	"fmt"
	"sync"
)

func search(tableau []int, element int, wg *sync.WaitGroup, results chan bool) {
	defer wg.Done() 
	for _, value := range tableau {
		if value == element {
			results <- true 
			return
		}
	}
	results <- false 
}

func main() {
	tableau := []int{2, 3, 5, 7, 11, 17, 19, 23, 29, 31, 37, 41, 43}
	elementCible := 19

	tailleSegment := 4

	var wg sync.WaitGroup       
	results := make(chan bool)  

	for i := 0; i < len(tableau); i += tailleSegment {
		wg.Add(1)
		fin := i + tailleSegment
		if fin > len(tableau) {
			fin = len(tableau)
		}
		go search(tableau[i:fin], elementCible, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	trouve := false
	for result := range results {
		if result {
			trouve = true
			break
		}
	}

	if trouve {
		fmt.Printf("L'élément %d a été trouvé dans le tableau.\n", elementCible)
	} else {
		fmt.Printf("L'élément %d n'a pas été trouvé dans le tableau.\n", elementCible)
	}

	fmt.Println("Recherche terminée.")
}
