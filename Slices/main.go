package main
import "fmt"

func main() {
	// Créer un tableau avec les nombres de 1 à 10
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arrayAvant:", array)


	// Créer un slice à partir du tableau
	numbers := array[:]

	// Ajouter 11 et 12 au slice numbers
	numbers = append(numbers, 11, 12)

	// Créer un sous-slice contenant les 5 premiers éléments
	firstNumbers := numbers[:5]

	// Modifier le deuxième élément du sous-slice
	firstNumbers[1] = 99

	// Afficher les deux slices
	fmt.Println("numbers:", numbers)
	fmt.Println("arrayLater:", array)
	fmt.Println("firstNumbers:", firstNumbers)
}