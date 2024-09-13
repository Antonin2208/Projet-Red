package projetred

import "fmt"

type Inventory struct {
	items map[string]int
}

func (inv *Inventory) removeInventory(item string, quantity int) {
	if _, exists := inv.items[item]; exists {
		if inv.items[item] > quantity {
			inv.items[item] -= quantity
		} else {
			delete(inv.items, item)
		}
	}
}

func (inv *Inventory) addInventory(item string, quantity int) {
	if _, exists := inv.items[item]; exists {
		inv.items[item] += quantity
	} else {
		inv.items[item] = quantity
	}
	fmt.Printf("Vous avez ajouté %d %s(s) à l'inventaire.\n", quantity, item)
}
