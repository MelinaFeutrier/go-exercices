package main

import "fmt"

type Vehicle struct {
	Brand string 
	Model string 
	Year  int    
}

func (v Vehicle) Description() {
	fmt.Printf("Ce véhicule est une %s %s de l'année %d.\n", v.Brand, v.Model, v.Year)
}

func main() {
	car := Vehicle{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2022,
	}

	car.Description()
}
