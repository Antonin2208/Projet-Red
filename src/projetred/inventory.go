package projetred

import (
	"fmt"

	"github.com/fatih/color"
)

// Fonction pour afficher l'inventaire du personnage
func ShowInventory(character *Character) {
	ClearScreen()
	color.Cyan("\nVotre Inventaire :")
	if len(character.Inventory) == 0 {
		color.Red("Votre inventaire est vide.")
		return
	}

	color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-")
	for i, item := range character.Inventory {
		color.Green(fmt.Sprintf("%d. %s (%s)", i+1, item.Name, item.Type))
	}
	color.White("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-")

	// Affiche les HP actuels
	color.Cyan(fmt.Sprintf("HP actuel: %d\n", character.HP))
}
