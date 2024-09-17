package projetred

import (
	"fmt"

	"github.com/fatih/color"
)

Fonction principale pour interagir avec l'inventaire
func InventoryMenu(character *Character) {

	ClearScreen()
	for {
		var choice int
		fmt.Print(`   
                     ____            __             _                                     __       __
                    / __ \  ____ _  / /_   ____    (_)        _      __  ____    _____   / /  ____/ /
                   / /_/ / / __  / / __/  / __ \  / /        | | /| / / / __ \  / ___/  / /  / __  / 
                  / _, _/ / /_/ / / /_   / /_/ / / /         | |/ |/ / / /_/ / / /     / /  / /_/ /  
                 /_/ |_|  \__,_/  \__/  / .___/ /_/          |__/|__/  \____/ /_/     /_/   \__,_/   
                                       /_/                                                           
	`)
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
 