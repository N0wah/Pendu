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
	motADeviner := "YNOVCAMPUS"
	lettresDevinees := make([]bool, len(motADeviner))
	lettreEntrer := ""
	tentativesRestantes := 6
	erreurs := 0

	fmt.Println(Cyan + "======================================================================================================================")
	fmt.Println("Bienvenue dans notre jeu, aujourd'hui vous devez sauver la vie de notre amie codeur pour qu'il vous donne un indice !")
	fmt.Println("Pour cela vous devez deviner le mot auquel le juge à penser pour lui éviter une pendaison.")
	fmt.Printf("          Le mot à deviner contient %d lettres.\n", len(motADeviner))
	fmt.Println("          À vous de jouer !!!")
	fmt.Println("======================================================================================================================" + Reset)

	for tentativesRestantes > 0 {
		fmt.Println("\nMot actuel:", motMasque(motADeviner, lettresDevinees))
		fmt.Printf("Tentatives restantes : %d\n", tentativesRestantes)
		afficherPendu(erreurs)
		fmt.Println("Lettres déja utilisées : " + lettreEntrer)

		fmt.Print("Entrez une lettre : ")
		lettre := lireLettre()

		if strings.Contains(lettreEntrer, lettre) {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Println(Red + "Veuillez entrer une lettre non utilisée." + Reset)
			continue
		}
		if lettre == "" || len(lettre) != 1  || !IsAlpha(lettre){
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Println(Red + "Veuillez entrer une seule lettre valide." + Reset)
			continue
		}
		lettreEntrer += lettre

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
			fmt.Println(Green+"Grâce à vous notre ami codeur est libre et vous donne l'indice qui est le chiffre", Red+"2"+Reset)
			break
		}
		fmt.Printf("\x1bc")
		fmt.Printf("\x1b[2J")
	}

	if tentativesRestantes == 0 {
		fmt.Println("   ____\n   |  |\n   O  |\n  /|\\ |\n  / \\ |\n      |\n=========")
		fmt.Println(Red + "\nOh non, vous avez épuisé toutes nos chances, notre amie codeur est mort." + Reset)
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

func IsAlpha(s string) bool {
	sRune := []rune(s)
	count := 0
	for i := range sRune {
		if sRune[i] >= 'A' && sRune[i] <= 'Z' || sRune[i] >= 'a' && sRune[i] <= 'z' {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
