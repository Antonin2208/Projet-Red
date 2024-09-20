package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Joueur struct {
	Nom         string
	Classe      string
	Niveau      int
	Experience  int
	SanteMax    int
	Sante       int
	Inventaire  map[string]int
	Pieces      int
	Competences []string
	Animal      string
	BoostDegats float64
	Energie     int
	Equipement  Equipement
}

type Monstre struct {
	Nom      string
	SanteMax int
	Sante    int
	Attaque  int
	XP       int
	ItemDrop string
}

type Equipement struct {
	Tete       string
	Torse      string
	Pieds      string
	Accessoire string
}

func EffacerConsole() {
	fmt.Print("\033[H\033[2J")
}

func MessageRapide(message string, vitesse int, nomCouleur string) {
	c := color.New(color.FgWhite)
	switch nomCouleur {
	case "vert":
		c = color.New(color.FgGreen)
	case "rouge":
		c = color.New(color.FgRed)
	case "bleu":
		c = color.New(color.FgBlue)
	case "cyan":
		c = color.New(color.FgCyan)
	case "jaune":
		c = color.New(color.FgYellow)
	}

	for _, char := range message {
		c.Print(string(char))
		time.Sleep(time.Duration(vitesse) * time.Millisecond)
	}
	fmt.Println()
}

func AfficherCadre(titre string, contenu []string) {
	largeur := 60
	fmt.Println("╔" + strings.Repeat("═", largeur-2) + "╗")
	fmt.Printf("║ %-*s ║\n", largeur-4, titre)
	fmt.Println("╠" + strings.Repeat("═", largeur-2) + "╣")
	for _, ligne := range contenu {
		fmt.Printf("║ %-*s ║\n", largeur-4, ligne)
	}
	fmt.Println("╚" + strings.Repeat("═", largeur-2) + "╝")
}

func (j *Joueur) Initialiser(nom, classe string) {
	j.Nom = nom
	j.Classe = classe
	j.Niveau = 1
	j.Experience = 0
	j.SanteMax = 100
	j.Sante = 100
	j.Inventaire = map[string]int{"Rhum": 2}
	j.Pieces = 100
	j.Competences = []string{"Coup de poing"}
	j.Animal = ""
	j.BoostDegats = 1.0
	j.Energie = 100
	j.Equipement = Equipement{
		Tete:       "Aucun chapeau",
		Torse:      "Aucune tunique",
		Pieds:      "Aucune sandale",
		Accessoire: "Aucun accessoire",
	}

	switch classe {
	case "Capitaine":
		j.Competences = append(j.Competences, "Tir de Mousquet")
	case "Navigateur":
		j.Competences = append(j.Competences, "Boussole Magique")
	case "Canonnier":
		j.Competences = append(j.Competences, "Tir de Canon")
	}
}

func (j *Joueur) AfficherInfo() {
	xpPourNiveauSuivant := j.Niveau * 100
	contenu := []string{
		fmt.Sprintf("Nom: %s", j.Nom),
		fmt.Sprintf("Classe: %s", j.Classe),
		fmt.Sprintf("Niveau: %d", j.Niveau),
		fmt.Sprintf("Expérience: %d / %d", j.Experience, xpPourNiveauSuivant),
		fmt.Sprintf("Santé: %d/%d", j.Sante, j.SanteMax),
		fmt.Sprintf("Pièces d'or: %d", j.Pieces),
		fmt.Sprintf("Inventaire: %v", j.Inventaire),
		fmt.Sprintf("Compétences: %v", j.Competences),
		fmt.Sprintf("Animal de compagnie: %s", j.Animal),
		fmt.Sprintf("Boost Dégâts: %.0f%%", (j.BoostDegats-1)*100),
		fmt.Sprintf("Équipement:\n  Tête: %s\n  Torse: %s\n  Pieds: %s\n", j.Equipement.Tete, j.Equipement.Torse, j.Equipement.Pieds),
	}
	AfficherCadre("Information du Pirate", contenu)
}

func (j *Joueur) UtiliserRhum() {
	if j.Inventaire["Rhum"] > 0 {
		j.Sante += 30
		if j.Sante > j.SanteMax {
			j.Sante = j.SanteMax
		}
		j.Inventaire["Rhum"]--
		fmt.Println("Vous avez bu une bouteille de rhum. Votre santé est maintenant :", j.Sante)
	} else {
		fmt.Println("Vous n'avez plus de rhum, moussaillon !")
	}
}

func (j *Joueur) ApprendreCompetence(competence string) {
	for _, c := range j.Competences {
		if c == competence {
			fmt.Printf("Vous connaissez déjà la compétence %s.\n", competence)
			return
		}
	}
	j.Competences = append(j.Competences, competence)
	fmt.Printf("Vous avez appris une nouvelle compétence : %s\n", competence)
}

func (j *Joueur) GagnerExperience(xp int) {
	j.Experience += xp
	xpPourNiveauSuivant := j.Niveau * 100
	if j.Experience >= xpPourNiveauSuivant {
		j.Niveau++
		j.Experience -= xpPourNiveauSuivant
		j.SanteMax += 10
		j.Sante = j.SanteMax
		fmt.Printf("Félicitations ! Vous êtes maintenant niveau %d !\n", j.Niveau)
	}
}

func EcranTitre() {
	EffacerConsole()
	titre := `
                             _                   _           _            _      
                        | |                 (_)         | |          (_)    
                        | |     __ _   _ __  _ _ __ __ _| |_ ___ _ __ _  ___
                        | |    / _  | | '_ \| | '__/ _  | __/ _ \ '__| |/ _ \
                        | |___| (_| | | |_) | | | | (_| | ||  __/ |  | |  __/
                        |______\__,_| | .__/|_|_|  \__,_|\__\___|_|  |_|\___|
                                  | |                                    
                                  |_|                                    
`
	color.Cyan(titre)
	MessageRapide("Une aventure épique sur les mers vous attend, moussaillon !", 50, "jaune")
	fmt.Println("\nAppuyez sur Entrée pour commencer votre voyage...")
	fmt.Println(`
 
 
                                  .                   .
                              _..-''"""\          _.--'"""\
                              |         L         |        \
                  _           / _.-.---.\.        / .-.----.\
                _/-|---     _/<"---'"""""\\'.    /-'--''"""""\
               |       \     |            L'.'-. |            L
               /_.-.---.L    |            \  \  '|            J'.
             _/--'""""  \    F            \L  \  |             L
               L      '. L  J  _.---.-"""-.\'. L_/ _.--|"""""--.\ '.
               |        \+. /-"__.--'""""   \ './'"---'""""""   \'. '.
               F   _____ \ 'F"        '.     \  \                L '.
              /.-'"_|---'"\ |           '    JL |                 L  '.'.
             <-'""         \|    _.-.------._ A J    _.-.-----'.--|   '.'.
              L         '. |/.-'"_.-'---'""\."| /-'"---'"""""   \'.\.  \ '.'.
              |  _.------\.<'"""            L\ L\                '.'\'. \  '.
         _.-'//'"--'"""   L\|       ________\ '.F     ___.-------._L \ '-\   \'.
        /___| F             F _.--'"_|-------L  /_.-'"_.-|-'"""""""\  L   L   '.'.
            | F  _.-'|"""""/'"-'"""          J <'"""                L J   |     '.'.
            |/-'-''/|""\ )-|\                 F \                   |  L .'"""'\""-\\_
             F'-'-'-('-')  | \                F  \                  |___'"""'.""'.-'"
------------/        -'---|  F               L   L             __     |"""""'-'"__________
          .'_         |    |__L          __  J__  |    _.--'""""   '".----'".'
         '""""""""""""|--._+--F _.-'""||"   """___/.-'"   ||-'"/""""" \_. .'
         J------------(___\__/'_____.--------'-------'""""""""           /
         '-.                                        _.__.__.__.____     J_.-._
    .''-._ (-'--'---.'--._'---._.-''-._.-'_.-''-._'               '-''-'
 
    `)

	fmt.Scanln()
}

func CreationPersonnage() Joueur {
	var j Joueur
	EffacerConsole()
	fmt.Println("Création du Pirate")
	fmt.Print("Entrez votre nom de pirate : ")
	fmt.Scanln(&j.Nom)

	contenu := []string{
		"1) Capitaine",
		"2) Navigateur",
		"3) Canonnier",
	}
	AfficherCadre("Choisissez votre classe de pirate", contenu)

	var choixClasse int
	fmt.Print("Votre choix (1-3) : ")
	fmt.Scanln(&choixClasse)

	switch choixClasse {
	case 1:
		j.Initialiser(j.Nom, "Capitaine")
	case 2:
		j.Initialiser(j.Nom, "Navigateur")
	case 3:
		j.Initialiser(j.Nom, "Canonnier")
	default:
		fmt.Println("Choix invalide. Par défaut : Capitaine.")
		j.Initialiser(j.Nom, "Capitaine")
	}

	return j
}

func MenuPrincipal(j *Joueur) {
	for {
		EffacerConsole()
		contenu := []string{
			"1. Afficher les informations du pirate",
			"2. Visiter la boutique du port",
			"3. Partir à l'aventure en mer",
			"4. Boire du rhum",
			"5. Visiter l'animalerie",
			"6. Forgeron",
			"7. Quitter le jeu",
		}
		AfficherCadre("Menu Principal", contenu)

		var choix int
		fmt.Print("Entrez votre choix (1-7) : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			j.AfficherInfo()
		case 2:
			BoutiquePort(j)
		case 3:
			AventureEnMer(j)
		case 4:
			j.UtiliserRhum()
		case 5:
			Animalerie(j)
		case 6:
			Forgeron(j)
		case 7:
			fmt.Println("Merci d'avoir joué à Pirate Adventure. À bientôt sur les mers, moussaillon !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Réessayez, moussaillon !")
		}

		fmt.Println("\nAppuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
}
func Forgeron(j *Joueur) {
	EffacerConsole()
	contenu := []string{
		"1. Fabriquer un Chapeau de pirate (15 pièces, 1 Plume de sirène , 1 Peau de requin)",
		"2. Fabriquer une Tunique de marin (15 pièces, 1 Dent, 1 Peau de requin)",
		"3. Fabriquer des Bottes de marin (15 pièces, 1 Peau de requin, 1 Bout de tentacule)",
		"4. Fabriquer un Cache-oeil (15 pièces, 1 Âme du Capitaine Phantom Alan)",
		"5. Retour au menu principal",
	}
	fmt.Printf("╔═══════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║ Menu forgeron                                                                                 ║\n")
	fmt.Printf("╣═══════════════════════════════════════════════════════════════════════════════════════════════╣\n")
	for _, i := range contenu {
		fmt.Println(fmt.Sprintf("%-95s ║", i))
	}
	fmt.Printf("╚═══════════════════════════════════════════════════════════════════════════════════════════════╝\n")

	var choix int
	fmt.Print("Entrez votre choix (1-5) : ")
	fmt.Scanln(&choix)

	if choix >= 1 && choix <= 4 {
		if j.Pieces < 15 {
			fmt.Println("Vous n'avez pas assez de pièces pour fabriquer cet équipement, moussaillon !")
			return
		}

		switch choix {
		case 1:
			if j.Inventaire["Plume de sirène"] < 1 || j.Inventaire["Peau de requin"] < 1 {
				fmt.Println("Vous n'avez pas les matériaux nécessaires pour fabriquer un Chapeau de pirate.")
			} else {
				j.Pieces -= 15
				j.Inventaire["Plume de sirène maléfique"]--
				j.Inventaire["Peau de requin"]--
				j.Inventaire["Chapeau de pirate"]++
				fmt.Println("Vous avez fabriqué un Chapeau de pirate !")
			}
		case 2:
			if j.Inventaire["Dent"] < 1 || j.Inventaire["Peau de requin"] < 1 {
				fmt.Println("Vous n'avez pas les matériaux nécessaires pour fabriquer une Tunique de marin.")
			} else {
				j.Pieces -= 15
				j.Inventaire["Dent"]--
				j.Inventaire["Peau de requin"]--
				j.Inventaire["Tunique de marin"]++
				fmt.Println("Vous avez fabriqué une Tunique de marin !")
			}
		case 3:
			if j.Inventaire["Peau de requin"] < 1 || j.Inventaire["Bout de tentacule"] < 1 {
				fmt.Println("Vous n'avez pas les matériaux nécessaires pour fabriquer des Bottes de marin.")
			} else {
				j.Pieces -= 15
				j.Inventaire["Peau de requin"]--
				j.Inventaire["Bout de tentacule"]--
				j.Inventaire["Bottes de marin"]++
				fmt.Println("Vous avez fabriqué des Bottes de marin !")
			}
		case 4:
			if j.Inventaire["Âme du Capitaine Phantom Alan"] < 1 {
				fmt.Println("Vous n'avez pas les matériaux nécessaires pour fabriquer un Cache-oeil.")
			} else {
				j.Pieces -= 15
				j.Inventaire["Âme du Capitaine Phantom Alan"]--
				j.Inventaire["Cache-oeil"]++
				fmt.Println("Vous avez fabriqué un Cache-oeil !")
			}
		}
	} else if choix == 5 {
		return
	} else {
		fmt.Println("Choix invalide, essayez à nouveau.")
	}
}

func BoutiquePort(j *Joueur) {
	EffacerConsole()
	contenu := []string{
		"1. Acheter une bouteille de rhum (10 pièces)",
		"2. Acheter un harpon (50 pièces)",
		"3. Améliorer l'épée (100 pièces)",
		"4. Augmenter l'inventaire de 4 places (15 pièces)",
		"5. Retour au menu principal",
	}
	AfficherCadre("Boutique du Port", contenu)

	var choix int
	fmt.Print("Entrez votre choix (1-4) : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if j.Pieces >= 10 {
			j.Pieces -= 10
			j.Inventaire["Rhum"]++
			fmt.Println("Vous avez acheté une bouteille de rhum !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 2:
		if j.Pieces >= 50 {
			j.Pieces -= 50
			j.ApprendreCompetence("Lancer de Harpon")
			fmt.Println("Vous avez acheté un harpon et appris Lancer de Harpon !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 3:
		if j.Pieces >= 100 {
			j.Pieces -= 100
			j.ApprendreCompetence("Coup d'Épée Puissant")
			fmt.Println("Vous avez amélioré votre épée et appris Coup d'Épée Puissant !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 4:
		if j.Pieces >= 15 {
			j.Inventaire["Nouvelle Place"] += 2
			j.Pieces -= 15
			fmt.Println("Vous avez augmenté votre inventaire de 2 places.")
		} else {
			fmt.Println("Vous n'avez pas assez de pièces.")
		}
	case 5:
		return
	default:
		fmt.Println("Choix invalide, moussaillon.")
	}
}

func InitialiserMonstre() Monstre {
	monstres := []Monstre{
		{"Kraken", 80, 80, 15, 50, "Bout de tentacule"},
		{"Sirène Maléfique", 60, 60, 12, 40, "Plume de sirène"},
		{"Requin Géant", 70, 70, 14, 45, "Peau de requin"},
		{"Pieuvre Géante", 75, 75, 13, 48, "Tentacule"},
		{"Capitaine Alan Phantom", 50, 50, 10, 35, "Âme spectrale"},
		{"Léviathan", 90, 90, 18, 60, "Dent"},
	}
	return monstres[rand.Intn(len(monstres))]
}

func (j *Joueur) CalculerEffetsEquipement() {
	j.SanteMax = 100
	j.BoostDegats = 1.0

	if j.Equipement.Tete == "Chapeau de pirate" {
		j.SanteMax += 5
	}
	if j.Equipement.Torse == "Tunique de marin" {
		j.SanteMax += 10
	}
	if j.Equipement.Pieds == "Sandales de pirate de l'aventurier" {
		j.SanteMax += 15
	}
	if j.Equipement.Tete == "Cache-œil" {
		j.BoostDegats += 0.05
	}

	if j.Sante > j.SanteMax {
		j.Sante = j.SanteMax
	}
}

func AventureEnMer(j *Joueur) {
	monstre := InitialiserMonstre()
	fmt.Printf("Un %s apparaît à l'horizon !\n", monstre.Nom)
	fmt.Println(`
 
||     *                            *                    ((   *  ||
||        *                 *                *            ~      ||
||                ___.                          *          *     ||
||       *    ___.\__|.__.           *                           ||
||            \__|. .| \_|.                                      ||
||            . X|___|___| .                         *           ||
||          .__/_||____ ||__.            *                /\     ||
||  *     .  |/|____ |_\|_ |/ _                          /  \    ||
||        \ _/ |_X__\|_  |\||~,~{                       /    \   ||
||         \/\ |/|    |_ |/:|X'{                   _ _/      \__||
||          \ \/ |___ |_\|_.|~~~                   /    . .. . ..||
||         _|X/\ |___\|_ :| |_.                  - .......... . .||
||         | __\_:____ |  ||o-|            ___/........ . . .. ..||
||         |/_-|-_|__ \|_ |/--|       ____/  . . .. . . .. ... . ||
|| ........:| -|- o-o\_:_\|o-/:....../...........................||
|| ._._._._._\=\====o==o==o=/:.._._._._._._._._._._._._._._._._._||
|| _._._._._._\_\ ._._._._.:._._._._._._._._._._._._._.P_._._._._||
|| ._._._._._._._._._._._._._._._._._._._._._._._._._._._._._._._||
||---------------------------------------------------------------||
 
    `)

	for j.Sante > 0 && monstre.Sante > 0 {
		contenu := []string{
			fmt.Sprintf("Votre Santé : %d/%d", j.Sante, j.SanteMax),
			fmt.Sprintf("Santé de %s : %d/%d", monstre.Nom, monstre.Sante, monstre.SanteMax),
		}
		AfficherCadre("Combat en Mer", contenu)

		fmt.Println("Choisissez votre action :")
		for i, competence := range j.Competences {
			fmt.Printf("%d. Utiliser %s\n", i+1, competence)
		}
		fmt.Printf("%d. Boire du rhum\n", len(j.Competences)+1)

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix > 0 && choix <= len(j.Competences) {
			degats := int(float64(rand.Intn(15)+10) * j.BoostDegats)
			if j.Competences[choix-1] == "Lancer de Harpon" {
				degats += 5
			}
			monstre.Sante -= degats
			fmt.Printf("Vous avez utilisé %s et infligé %d dégâts à %s !\n", j.Competences[choix-1], degats, monstre.Nom)
		} else if choix == len(j.Competences)+1 {
			j.UtiliserRhum()
		} else {
			fmt.Println("Choix invalide. Vous avez manqué votre tour !")
		}

		if monstre.Sante <= 0 {
			fmt.Printf("Vous avez vaincu %s !\n", monstre.Nom)
			butin := rand.Intn(50) + 30
			j.Pieces += butin
			fmt.Printf("Vous avez trouvé %d pièces d'or !\n", butin)

			item := monstre.ItemDrop
			j.Inventaire[item]++
			fmt.Printf("Vous avez obtenu un(e) %s !\n", item)

			j.GagnerExperience(monstre.XP)
			return
		}

		degatsMonstre := monstre.Attaque + rand.Intn(5)
		j.Sante -= degatsMonstre
		fmt.Printf("%s vous attaque et inflige %d dégâts !\n", monstre.Nom, degatsMonstre)
	}

	fmt.Println("Vous avez été vaincu...")
}

func Animalerie(j *Joueur) {
	EffacerConsole()
	fmt.Println("Bienvenue à l'animalerie !")
	fmt.Println("1. Adopter Keir l'orque (15% boost de dégâts, 25 pièces)")
	fmt.Println("2. Adopter Cyril l'hippocampes (15% vie, 50 pièces)")
	fmt.Println("3. Adopter marie la sirène (beauté , 999 pièces)")
	fmt.Println("4. Adopter Maxime la raie (boost market , 10 pièces)")
	fmt.Println("5. Retour")
	fmt.Println(`
                \:.|"._
       /\/;.:':::;;;._
            <  .'     ':::;(
             < ' _      '::;>
              \ (9)  _  :::;(
              |     / \   ::;>
              |    /  |    :;(
              |   (  <=-  .::;>
              (  a) )=-  .::;(
               '-' <=-  .::;>
                  )==- ::::(  ,
                 <==-  :::(,-'(
                 )=-   '::  _.->
                <==-    ':.' _(
                 <==-    .:'_ (
                  )==- .::'  '->
                   <=- .:;('.(
                    )  ':;>  
               .-.  <    :;(
             <.':\  )    :;>
            < :/<_/  <  .:;>
            < '---'  .::(
             <       .:;>'
              -..:::-'                  
`)
	fmt.Println(`
 
        (\.-./)
                /     \
              .'   :   '.
         _.-''     '     ''-._
      .-'          :          '-.
    ,'_.._         .         _.._',
    ''    ''-.     '     .-''    ''
              '.   :   .'
                \_. ._/
          \       |^|
           |      | ;
           \'.___.' /
            '-....-'
`)

	fmt.Println(`
⠀⠀⠀⠀⠀⠀⣠⣴⣶⣆⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣷⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠐⢻⣿⣿⣿⣿⣿⣿⡅⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢿⡿⠿⣿⣿⣿⡿⣯⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣀⣼⣿⡿⢿⣮⣷⣄⡀⠀⠀
⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄
⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⠿⢛⣿⣿⠟⠁
⠀⠀⢀⣤⣤⣭⡛⠛⠉⢉⣤⣶⠟⠛⠁⠀⠀
⣤⣾⣿⣿⣿⣿⣿⣿⡿⠛⠁⠀⠀⠀⠀⠀⠀
⠀⠙⠛⠛⠛⣻⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣿⣿⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠈⢿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠻⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
   
   
    `)

	fmt.Println(`
   
                                  _
                                     / |
                                    /  |
                                   /   |
                                  /    |
                ,____.-----------'     ".______,                         ._
           ._--"                                "-----.______,          / /
         ,/__o_~##mm                                          "-----__./  |
         ==--    ~~                            ____                    \_ |
          "~###**~~~|   \                    ./####"            _________(
               ~~~##|    |~~~      _________/###/                /###\.   |
                    (    |##################~__________--------~~~~~~~~\  |
                     \   |  ~~###########~~~~~~                         \._\
                      \__\
   
   
   
    `)

	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if j.Pieces >= 10 {
			j.Pieces -= 10
			j.Animal = "Keir l'orque"
			j.BoostDegats *= 1.15
			fmt.Println("Vous avez adopté Keir l'orque !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 2:
		if j.Pieces >= 10 {
			j.Pieces -= 10
			j.Animal = "Cyril l'hippocampes"
			j.SanteMax = int(float64(j.SanteMax) * 1.15)
			j.Sante = j.SanteMax
			fmt.Println("Vous avez adopté Cyril l'hippocampes !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 3:
		if j.Pieces >= 999 {
			j.Pieces -= 999
			j.Animal = "Marie la sirène"
			fmt.Println("Vous avez adopté Marie la sirène !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 4:
		if j.Pieces >= 10 {
			j.Pieces -= 10
			j.Animal = "Maxime la raie"
			fmt.Println("Vous avez adopté Maxime la raie !")
		} else {
			fmt.Println("Pas assez de pièces, moussaillon.")
		}
	case 5:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	EcranTitre()
	joueur := CreationPersonnage()
	MenuPrincipal(&joueur)
}
