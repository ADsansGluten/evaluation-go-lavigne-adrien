package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


func afficherMenu() int{   // affiche de base et recup le choix en int

 var choix int 

	println("=== DISTRIBUTEUR ===\n")
	println("1 - Eau       : 1 $")
	println("2 - Soda      : 2 $")
	println("3 - Café      : 2 $")
	println("4 - Chocolat  : 3 $")
	println("0 - Quitter\n")

	println("Dis moi le numéro et vite hein !\n")

	print("Je choisie le : ")
	fmt.Scan(&choix)   

	switch choix {   // on essaye d'innover muehehehe
		case 1: 
			println("Pfff tu prends de l'eau au distributeur..")
			println("Ca sera 1$ pour toi.")
		case 2: 
			println("Ouai ouai c'est ca, fais gaffe à toi \nSi tu prends pas un Soda !!")
			println("Envoie un 2$.")
		case 3: 
			println("Por la manana Cafeeee, por la tarde ron !!")
			println("Dos dollares pour le BadBunny !")
		case 4: 
			println("T'as crue j'étais un disney pour un Chocolat ??")
			println("Nois-toi... mais avant 3$ stp.")
		case 0: 
			println("Bon casse toi alors...")
		default :                             // autres choix que 12340
			println("! SAISIE INVALIDE !")
	}
	return choix 
}

func quitter(choix int) bool{  // me fait quitter le programme, retourne booléen
	if choix == 0 {
		return true
	}
	return false
}

func obtenirPrix(choix int) int {  // renvoie mon prix en int
	switch choix {
		case 1 : 
			return 1
		case 2, 3: 
			return 2
		case 4 : 
			return 3
		default :     // ca me sert pour n'importe quels choix autre que 1234
			return 0
	}
}


func main() {  // organise tout, genre il redirige
	for {
		choix := afficherMenu()  //retourne un int

			if quitter(choix) {  // recois un int et retourne un bool
				time.Sleep(3 * time.Second)
				clearTerminal()  
				break  // j'aurais mis continue mais t'as dis fallait quitter le programme
			}
			
		prix := obtenirPrix(choix)  // recois un int et retourne un int

			if prix > 0 {
				fmt.Println("Prix :", prix, "$")

				monnaie := obtenirThune(prix)
					fmt.Println("Votre monnaie :", monnaie, "$")
				
				fmt.Println("\nLe menu va revenir dans 5 secondes...")
			  	time.Sleep(5 * time.Second)

				clearTerminal()
			}
		fmt.Println()
	}
}	

func obtenirThune(prix int) int{   // calcule la monnaie à rendre
	var Montant int 
	total := 0

	for total < prix {
		fmt.Print("Montant inséré : ")
		fmt.Scan(&Montant)

		total += Montant

		if total < prix {    // vérifie si l'argent inséré est suffisant
			fmt.Println("EH il manque là !")
			fmt.Println("Donne", prix - total, "$.")
		}

	}

	Thune := total - prix

	fmt.Println()
	fmt.Println("ARRIIIBAAAA !")  //ref à SDM

	return Thune
}

func clearTerminal() {  // le prime du prime pour faire propre 
	cmd := exec.Command("cmd", "/c", "cls") //lance une commande comme si c'est moi qui écrivait
	cmd.Stdout = os.Stdout  // relie la sortie au terminal, genre ca fait le lien
	cmd.Run()
}
