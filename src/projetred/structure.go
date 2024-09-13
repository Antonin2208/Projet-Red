package projetred

// Définition de la structure d'un personnage (Character)
type Character struct {
	Name      string
	Class     string
	HP        int
	Inventory []Item // L'inventaire contient une liste d'objets
}

// Définition de la structure d'un objet (Item)
type Item struct {
	Name    string
	Type    string // Par exemple "Arme", "Potion", "Armure"
	HPBonus int    // Bonus de HP si applicable
}
