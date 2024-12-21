package main

import "fmt"

func main() {
    m := createSpiral(3)
    for _, row := range m {
        fmt.Println(row)
    }
}

func createSpiral(n int) [][]int {
	//Retourner un tableau vide si N < 1 ou si N n'est pas un nombre entier
    if n < 1 {
        return [][]int{}
    }

    matrice := make([][]int, n)

    for i := 0; i < n; i++ {
        matrice[i] = make([]int, n)
    }

    haut := 0
    bas := n - 1
    gauche := 0
    droite := n - 1
    num := 1

    for haut <= bas && gauche <= droite {
        for i := gauche; i <= droite; i++ {
            matrice[haut][i] = num
            num++
        }
        haut++

        for i := haut; i <= bas; i++ {
            matrice[i][droite] = num
            num++
        }
        droite--

        if haut <= bas {
            for i := droite; i >= gauche; i-- {
                matrice[bas][i] = num
                num++
            }
            bas--
        }

        if gauche <= droite {
            for i := bas; i >= haut; i-- {
                matrice[i][gauche] = num
                num++
            }
            gauche++
        }
    }

    return matrice
}
