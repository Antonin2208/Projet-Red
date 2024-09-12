package projetred

import "fmt"

func ShowMenu() string {
	var name string

	fmt.Println("Bienvenue dans notre jeux !")
	fmt.Println("Veuillez entrer votre nom :")
	fmt.Scanln(&name)

	return name
}
