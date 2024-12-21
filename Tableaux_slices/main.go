package main

import "fmt"

func main(){
	// Call the function
	printDays()
}

func printDays(){
	//declare un tableau

	days := [7]string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}
	//on initialise le tableau 
	
	// Parcourir le tableau et afficher les jours
	//on affiche les jours de la semaine
	for i := 0; i < len(days);  i++ {
		fmt.Println(days[i])
	}

	// Modifier "Mercredi" en "Woden's Day"
	days[2] = "Woden's Day"	
	for i := 0; i < len(days);  i++ {
		fmt.Println(days[i])
	}
	
	//fmt.Println("hello")
	
}

