package main

import "fmt"

func main() {
	reponse := "Salut"
	runeReponse := []rune(reponse)
	affichage := ""
	for i := 0; i != len(reponse); i++ {
		affichage = affichage + "_"
	}
	runeAffichage := []rune(affichage)
	trouve := 0
	coups := 0
	for trouve != len(reponse) {
		var rep string
		fmt.Println(string(runeAffichage))
		fmt.Scanln(&rep)
		for i := range runeReponse {
			if rep == string(runeReponse[i]) {
				runeAffichage[i] = runeReponse[i]
				trouve++
			}
		}
		coups++
	}
	fmt.Println(string(runeAffichage))
	fmt.Printf("Tu as trouvé la réponse en ")
	fmt.Println(coups)
}
