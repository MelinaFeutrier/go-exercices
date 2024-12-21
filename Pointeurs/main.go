package main

import "fmt"

func inverser(a, b *int) {
    *a = *a + *b 
    *b = *a - *b 
    *a = *a - *b 
}

func main() {
    x := 5
    y := 10

    fmt.Println("Avant: x =", x, ", y =", y)

    inverser(&x, &y)

    fmt.Println("Après: x =", x, ", y =", y)
}
