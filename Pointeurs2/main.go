package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func ModifyPerson(p *Person) {
    p.Name = "Coucou" 
    p.Age += 100            
}

func main() {
    person := Person{Name: "Alice", Age: 25}

    fmt.Println("Avant modification :", person)

    ModifyPerson(&person)

    fmt.Println("Après modification :", person)
}
