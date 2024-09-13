package projetred

import "fmt"

// ShowMenu affiche un menu d'accueil et demande le nom du joueur
func ShowMenu() string {
	var name string

	fmt.Println(`

`)
	fmt.Println("Bienvenue dans notre jeu !")
	fmt.Print("Veuillez entrer votre nom : ")
	fmt.Scanln(&name)

	// Confirmer l'entrée du joueur
	fmt.Printf("Merci, %s ! Préparez-vous pour l'aventure !\n", name)

	return name
}
