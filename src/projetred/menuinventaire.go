package projetred

import (
	"fmt"

	"github.com/fatih/color"
)

// Fonction principale pour interagir avec l'inventaire
func InventoryMenu(character *Character) {
	ClearScreen()
	for {
		var choice int
		color.Cyan("\nQue voulez-vous faire ?")
		color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- ")
		color.Green("1. Afficher l'inventaire")
		color.Red("2. Quitter")
		color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_- ")
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ShowInventory(character)
		case 2:
			color.Red("Vous quittez le menu de l'inventaire.")
			return
		default:
			color.Red("Choix invalide. Veuillez essayer à nouveau.")
		}
	}
}
