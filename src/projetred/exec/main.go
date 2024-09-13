package main

import (
	"fmt"
	"projetred"

	"github.com/fatih/color"
)

func MenuPrincipal() {
	projetred.ClearScreen()
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

		color.White(" ")
		fmt.Println()
		color.Cyan("1. Stats")
		fmt.Println()
		color.Green("2. Inventaire")
		fmt.Println()
		color.Yellow("3. Marchand")
		fmt.Println()
		color.Red("4. Quitter")
		fmt.Println()
		color.White("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂")
		fmt.Println()
		fmt.Print("Entrez votre choix (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Vous avez choisi d'afficher un message.")
		case 2:

		case 3:
			projetred.MenuMarchand()
		case 4:
			projetred.ClearScreen()
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Choix invalide. Veuillez entrer un nombre entre 1 et 3.")

		}
	}
}

func main() {
	character := projetred.CreateCharacter()
	projetred.InventoryMenu(&character)
	MenuPrincipal()
}
