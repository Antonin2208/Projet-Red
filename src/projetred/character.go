package projetred

import (
	"fmt"

	"github.com/fatih/color"
)

var classItems = map[string]Item{
	"Guerrier": {Name: "Épée", Type: "Arme", HPBonus: 0},
	"Mage":     {Name: "Baguette", Type: "Arme", HPBonus: 0},
	"Archer":   {Name: "Arc", Type: "Arme", HPBonus: 0},
}

func CreateCharacter() Character {
	var name, class string
	var hp int

	fmt.Print(`   
                                   ____            __             _                                     __       __
                                  / __ \  ____ _  / /_   ____    (_)        _      __  ____    _____   / /  ____/ /
                                 / /_/ / / __  / / __/  / __ \  / /        | | /| / / / __ \  / ___/  / /  / __  / 
                                / _, _/ / /_/ / / /_   / /_/ / / /         | |/ |/ / / /_/ / / /     / /  / /_/ /  
                               /_/ |_|  \__,_/  \__/  / .___/ /_/          |__/|__/  \____/ /_/     /_/   \__,_/   
                                                     /_/                                                           
	`)

	fmt.Println()
	fmt.Println()
	fmt.Print("						  		lâche ton blaze poto :")

	fmt.Scanln(&name)
	fmt.Println()

	fmt.Println("					  		 	Choisissez une classe ")
	fmt.Println("								1. Ratpi vilain 🔫")
	fmt.Println("								2. Magicien tah gandalf 🧙 ")
	fmt.Println("								3. Archer ")
	fmt.Println()
	fmt.Print("								Votre choix (1-3) : ")
	fmt.Scanln(&class)

	switch class {
	case "1":
		class = "Guerrier"
		hp = 150
	case "2":
		class = "Mage"
		hp = 100
	case "3":
		class = "Archer"
		hp = 120
	default:
		color.Red("Choix invalide. Vous serez un Aventurier par défaut.")
		class = "Aventurier"
		hp = 110
	}

	initialItem := classItems[class]
	inventory := []Item{initialItem}

	return Character{Name: name, Class: class, HP: hp, Inventory: inventory}
}
