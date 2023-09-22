package main

import "fmt"

func main() {
	fmt.Printf("\x1b[2J")
	reponse := "Salut"
	runeReponse := []rune(reponse)
	affichage := ""
	for i := 0; i != len(reponse); i++ {
		affichage = affichage + "_"
	}
	runeAffichage := []rune(affichage)
	trouve := 0
	coups := 0
	erreur := 0
	for trouve != len(reponse) {
		var rep string
		fmt.Println(string(runeAffichage))
		fmt.Scanln(&rep)
		trouveReg := trouve
		for i := range runeReponse {
			if rep == string(runeReponse[i]) {
				runeAffichage[i] = runeReponse[i]
				trouve++
			}
		}
		if trouveReg == trouve {
			erreur++
		}
		coups++
		if erreur == 1 {
			fmt.Printf("\x1b[2J")
			Dessin1()
		} else if erreur == 2 {
			fmt.Printf("\x1b[2J")
			Dessin2()
		} else if erreur == 3 {
			fmt.Printf("\x1b[2J")
			Dessin3()
		} else if erreur == 4 {
			fmt.Printf("\x1b[2J")
			Dessin4()
		} else if erreur == 5 {
			fmt.Printf("\x1b[2J")
			Dessin5()
			break
		}
	}
	fmt.Println(string(runeAffichage))
	if string(runeAffichage) == string(runeReponse) {
		fmt.Printf("Tu as trouvé la réponse en ")
		fmt.Println(coups)
	} else {
		fmt.Println("Tu n'as pas trouvé, dommage !")
	}
}

func Dessin1() {
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("        ")
	fmt.Println("________")
}

func Dessin2() {

	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("________")
}
func Dessin3() {
	fmt.Println(" _______|")
	fmt.Println("        |")
	fmt.Println("        |")
	fmt.Println("        |")
	fmt.Println("        |")
	fmt.Println("        |")
	fmt.Println("        |")
	fmt.Println("_________")
}

func Dessin4() {
	fmt.Println("_______|")
	fmt.Println("   |   |")
	fmt.Println("   0   |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("       |")
	fmt.Println("________")
}

func Dessin5() {
	fmt.Println(" _______|")
	fmt.Println("    |   |")
	fmt.Println("    0   |")
	fmt.Println("   \\|/  |")
	fmt.Println("    |   |")
	fmt.Println("   / \\  |")
	fmt.Println("        |")
	fmt.Println("_________")
}
