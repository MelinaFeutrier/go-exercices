package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go est génial. Go est rapide et Go est simple. GÉNIAL, non ?"

	wordFrequency := wordsCount(text)

	fmt.Println("Fréquence des mots :")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func wordsCount(texte string) map[string]int {

	wordCount := make(map[string]int)

	texte = strings.ToLower(texte)
    texte = strings.ReplaceAll(texte, ".", " ")
	texte = strings.ReplaceAll(texte, ",", " ")
	texte = strings.ReplaceAll(texte, "?", " ")
    texte = strings.ReplaceAll(texte, "!", " ")
    texte = strings.ReplaceAll(texte, "'", " ")

	
	words := strings.Fields(texte)

    for i := 0; i < len(words);  i++ {
		wordCount[words[i]]++
	} 

	return wordCount
}
