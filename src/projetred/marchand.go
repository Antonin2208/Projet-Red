package projetred

import (
	"fmt"

	"github.com/fatih/color"
)

func MenuMarchand() {
	for {
		var choice int

		fmt.Println("Tu veux un 10 balle ? non je rigole viens la je vais te montrer mes meilleur plavon")
		color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- ")
		fmt.Println()
		color.Cyan("1.  ")
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
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide. Veuillez entrer un nombre entre 1 et 3.")

		}
	}
}
