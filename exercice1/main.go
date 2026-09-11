package main

import (
	"fmt"
	"go/printer"
)

func quitter(choix int) bool{
	if choix == 0 {
		println("Bon casse toi alors...")
		return true
	}
	return 

}


func afficherMenu(

	println("=== DISTRIBUTEUR ===\n")
	println("1 - Eau       : 1 €")
	println("2 - Soda      : 2 €")
	println("3 - Café      : 2 €")
	println("4 - Chocolat  : 3 €")
	println("0 - Quitter\n")

	println("Dis moi le numéro et vite hein !\n")

	println("Je choisie le : ")
	Scan(&choix)

	switch choix {
		case 1: 
			println("Pfff tu prends de l'eau au distributeur..")
			println("Ca sera 1$ pour toi.")
		case 2: 
			println("Ouai ouai c'est ca, fais gaffe à toi \n Si tu prends pas un Soda !!")
			println("Envoie un 2$.")
		case 3: 
			println("Por la manana Cafeeee, por la tarde ron !!")
			println("Tres dollars pour le BadBunny !")
		case 4: 
			println("T'as crue j'étais un disney pour un Chocolat ??")
			println("Nois-toi... mais avant 3$ stp.")
		case 0: 
			println("Bon casse toi alors...")
		default : 
			println("!SAISIE INVALIDE!")
	}
	
)
	if choix < 0 || > 4 {
		return true
	}
	return false

	if choix = 0 {
		println 

		
	}


)

func obtenirPrix(choix int) int(
if choix < 1 || > 4 {
		return true
	}
	return false



)


func donneargent(prix int) int(
if 


)



fmt.Println("Choisie une boisson !!")
fmt.Scan(boisson)

func main (boisson int) int (




)