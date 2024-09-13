package projetred

import (
	"fmt"
)

func CreateCharacter() Character {
	var name, class string

	// Demande du nom du personnage
	fmt.Println(`
	
	
	`)
	fmt.Scanln(&name)

	// Sélection de la classe du personnage
	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Archer")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&class)

	switch class {
	case "1":
		class = "Guerrier"
	case "2":
		class = "Mage"
	case "3":
		class = "Archer"
	default:
		fmt.Println("Choix invalide, vous serez un Aventurier par défaut.")
		class = "Aventurier"
	}

	// Retourne un nouveau personnage
	return Character{Name: name, Class: class}
}
