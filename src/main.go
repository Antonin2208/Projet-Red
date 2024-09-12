package main

import (
	"fmt"
)

func ScanInput() {
	for {
		var choice int

		fmt.Println("Menu:")
		fmt.Println("=====================================================================")
		fmt.Println("1. Afficher un message")
		fmt.Println("2. Entrer un nombre")
		fmt.Println("3. Quitter")
		fmt.Println("=====================================================================")
		fmt.Print("Entrez votre choix (1-3): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Vous avez choisi d'afficher un message.")
		case 2:
			var num int
			fmt.Print("Tapez un nombre: ")
			fmt.Scan(&num)
			fmt.Println("Votre nombre est:", num)
		case 3:
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide. Veuillez entrer un nombre entre 1 et 3.")
		}
	}
}

func main() {
	ScanInput()
}
