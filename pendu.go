package main

import (
	"fmt"
	"strings"
)

// Couleurs ANSI
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func main() {
	motADeviner := "GOLANG"
	lettresDevinees := make([]bool, len(motADeviner))
	tentativesRestantes := 6
	erreurs := 0

	fmt.Println(Cyan + "===================================")
	fmt.Println("Bienvenue dans le jeu du pendu !")
	fmt.Printf("Le mot à deviner contient %d lettres.\n", len(motADeviner))
	fmt.Println("===================================" + Reset)

	for tentativesRestantes > 0 {
		fmt.Println("\nMot actuel:", motMasque(motADeviner, lettresDevinees))
		fmt.Printf("Tentatives restantes : %d\n", tentativesRestantes)
		afficherPendu(erreurs)

		fmt.Print("Entrez une lettre : ")
		lettre := lireLettre()
		if lettre == "" || len(lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre valide.")
			continue
		}

		lettre = strings.ToUpper(lettre)

		lettreRatée := true
		for i, c := range motADeviner {
			if lettre[0] == byte(c) {
				lettresDevinees[i] = true
				lettreRatée = false
			}
		}

		if lettreRatée {
			tentativesRestantes--
			erreurs++
		}

		if motMasque(motADeviner, lettresDevinees) == motADeviner {
			fmt.Printf(Green+"\nFélicitations, vous avez deviné le mot : %s\n"+Reset, motADeviner)
			break
		}

		fmt.Printf("\x1bc")
		fmt.Printf("\x1b[2J")
	}

	if tentativesRestantes == 0 {
		fmt.Println("   ____\n   |  |\n   O  |\n  /|\\ |\n  / \\ |\n      |\n=========")
		fmt.Printf(Red+"\nDésolé, vous avez épuisé toutes vos tentatives. Le mot était : %s\n"+Reset, motADeviner)
	}
}
func motMasque(mot string, lettresDevinees []bool) string {
	motMasque := ""
	for i, c := range mot {
		if lettresDevinees[i] {
			motMasque += string(c)
		} else {
			motMasque += "_"
		}
	}
	return motMasque
}

func lireLettre() string {
	var lettre string
	_, err := fmt.Scanln(&lettre)
	if err != nil {
		return ""
	}
	return lettre
}

func afficherPendu(erreurs int) {
	pendu := []string{
		"",
		"   ____\n   |  |\n      |\n      |\n      |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n      |\n      |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n   |  |\n      |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n  /|  |\n      |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n  /|\\ |\n      |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n  /|\\ |\n  /   |\n      |\n=========",
		"   ____\n   |  |\n   O  |\n  /|\\ |\n  / \\ |\n      |\n=========",
	}
	fmt.Println(pendu[erreurs])
}
