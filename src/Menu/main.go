package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
}

// Structure pour représenter un objet
type Item struct {
	Name        string
	Description string
}

// Structure pour représenter un inventaire
type Inventory struct {
	Items map[string]Item // Utilisation d'une map pour un accès rapide
}

// Fonction pour créer un nouvel objet
func NewItem(name, description string) Item {
	return Item{Name: name, Description: description}
}

// Fonction pour créer un nouvel inventaire
func NewInventory() Inventory {
	return Inventory{Items: make(map[string]Item)}
}

// Méthode pour ajouter un objet à l'inventaire
func (inv *Inventory) AddItem(item Item) {
	inv.Items[item.Name] = item
	fmt.Printf("L'objet '%s' a été ajouté à l'inventaire.\n", item.Name)
}

// Méthode pour retirer un objet de l'inventaire
func (inv *Inventory) RemoveItem(itemName string) {
	if _, exists := inv.Items[itemName]; exists {
		delete(inv.Items, itemName)
		fmt.Printf("L'objet '%s' a été retiré de l'inventaire.\n", itemName)
	} else {
		fmt.Printf("L'objet '%s' n'existe pas dans l'inventaire.\n", itemName)
	}
}

// Méthode pour afficher le contenu de l'inventaire
func (inv *Inventory) Display() {
	fmt.Println("Inventaire :")
	if len(inv.Items) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	for _, item := range inv.Items {
		fmt.Printf("Nom: %s, Description: %s\n", item.Name, item.Description)
	}
}

func main() {
	// Créer un inventaire
	clearScreen()
	inventory := NewInventory()

	for {
		fmt.Println("\nMenu principal :")
		fmt.Println("1. Ajouter un objet à l'inventaire")
		fmt.Println("2. Retirer un objet de l'inventaire")
		fmt.Println("3. Afficher le contenu de l'inventaire")
		fmt.Println("4. Quitter")

		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Entrée invalide. Veuillez entrer une lettre.")
			continue
		}

		switch choice {
		case 1:
			var name, description string
			fmt.Print("Entrez le nom de l'objet : ")
			fmt.Scan(&name)
			fmt.Print("Entrez la description de l'objet : ")
			fmt.Scan(&description)
			item := NewItem(name, description)
			inventory.AddItem(item)

		case 2:
			var name string
			fmt.Print("Entrez le nom de l'objet à retirer : ")
			fmt.Scan(&name)
			inventory.RemoveItem(name)

		case 3:
			inventory.Display()

		case 4:
			fmt.Println("Quitter le programme.")
			return

		default:
			fmt.Println("Choix non valide. Essayez encore.")
		}
	}
}
