package projetred

import "fmt"

// ShowMenu affiche un menu d'accueil et demande le nom du joueur
func ShowMenu() string {
	var name string

	fmt.Println(`

`)
	fmt.Println("Ouais mon gars la t'arrive dans le game du game tes pas prêt")
	fmt.Print("lache ton blaze fréro : ")
	fmt.Scanln(&name)

	// Confirmer l'entrée du joueur
	fmt.Printf("Merci, %s ! Préparez-vous pour l'aventure !\n", name)

	return name
}
