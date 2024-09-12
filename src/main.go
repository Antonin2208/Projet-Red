package main

import (
	"fmt"
	"projetred/projetred"

	"github.com/fatih/color"
)

func MenuPrincipal() {
	for {
		var choice int

		fmt.Println("Menu:")
		color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- ")
		fmt.Println()
		color.Cyan("1. Stats")
		fmt.Println()
		color.Green("2. Inventaire")
		fmt.Println()
		color.Yellow("3. Marchand")
		fmt.Println()
		color.Red("4. Quitter")
		fmt.Println()
		color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- ")
		fmt.Println()
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
			projetred.MenuMarchand()
		case 4:
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide. Veuillez entrer un nombre entre 1 et 3.")

		}
	}
}

func main() {
	MenuPrincipal()
}
