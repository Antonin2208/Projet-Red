package projetred

import (
	"fmt"
	"github.com/fatih/color"
)

// Liste des objets disponibles par classe
var classItems = map[string]Item{
	"Guerrier": {Name: "Épée", Type: "Arme", HPBonus: 0},
	"Mage":     {Name: "Baguette", Type: "Arme", HPBonus: 0},
	"Archer":   {Name: "Arc", Type: "Arme", HPBonus: 0},
}

// Initialisation d'un personnage avec nom, classe, HP et inventaire
func CreateCharacter() Character {
	var name, class string
	var hp int

	// Demande du nom du personnage
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	// Choix de la classe
	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Archer")
	fmt.Print("Votre choix (1-3) : ")
	fmt.Scanln(&class)

	switch class {
	case "1":
		class = "Guerrier"
		hp = 150 // HP pour Guerrier
	case "2":
		class = "Mage"
		hp = 100 // HP pour Mage
	case "3":
		class = "Archer"
		hp = 120 // HP pour Archer
	default:
		color.Red("Choix invalide. Vous serez un Aventurier par défaut.")
		class = "Aventurier"
		hp = 110
	}

	// Ajoute l'item spécifique à la classe dans l'inventaire
	initialItem := classItems[class]
	inventory := []Item{initialItem}

	// Retourne le personnage initialisé
	return Character{Name: name, Class: class, HP: hp, Inventory: inventory}
}
