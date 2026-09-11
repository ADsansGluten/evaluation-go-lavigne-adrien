# Évaluation pratique — Fondamentaux de Golang

**Langage :** Go  
**Barème :** /20  
**Dépôt :** Git obligatoire

---

# 1. Objectifs de l'évaluation

Cette évaluation permet de vérifier votre capacité à :

- créer et exécuter un programme Go ;
- déclarer et utiliser des variables ;
- récupérer des informations saisies par l'utilisateur ;
- utiliser des conditions `if / else` ;
- utiliser des boucles `for` ;
- utiliser les opérateurs de comparaison et `%` ;
- construire une logique de programme ;
- créer et utiliser des fonctions simples ;
- utiliser Git pour sauvegarder et organiser son travail.

**Important :** l'objectif n'est pas uniquement d'obtenir un programme qui fonctionne. Le code doit être compréhensible et correctement organisé.

---

# 2. Contraintes Git

## 2.1 Création du dépôt

Vous devez créer un dépôt Git dont le nom respecte exactement le format suivant :

```text
evaluation-go-nom-prenom
```

### Exemple

Pour l'étudiant Jean Dupont :

```text
evaluation-go-jean-dupont
```

Le dépôt doit être **public**.

---

## 2.2 Structure du dépôt

Votre dépôt doit contenir au minimum :

```text
evaluation-go-nom-prenom/
│
├── README.md
│
├── exercice1/
│   └── main.go
│
└── exercice2/
   └── main.go
```

---

## 2.3 README.md

Le fichier `README.md` doit contenir :

```text
# Évaluation Go

Nom :
Prénom :

## Exercices

### Exercice 1
Description rapide de votre programme.

### Exercice 2
Description rapide de votre programme.
```

---

# 3. Que les exercices commencent !!

## Exercice 1 — Le distributeur de boissons

**Difficulté : ⭐⭐⭐**  
**Notation : 8pts**

Vous devez créer un programme simulant un petit distributeur de boissons.

Le distributeur propose 4 boissons :

```text
1 - Eau       : 1 €
2 - Soda      : 2 €
3 - Café      : 2 €
4 - Chocolat  : 3 €
0 - Quitter
```

### Fonctionnement

Le programme affiche le menu et demande à l'utilisateur de choisir une boisson.

Exemple :

```text
=== DISTRIBUTEUR ===

1 - Eau       : 1 €
2 - Soda      : 2 €
3 - Café      : 2 €
4 - Chocolat  : 3 €
0 - Quitter

Votre choix : 2

Vous avez choisi : Soda
Prix : 2 €
```

Le programme demande ensuite combien d'argent l'utilisateur insère :

```text
Montant inséré : 5 €

Merci !
Votre monnaie : 3 €
```

---

### Contraintes

Vous devez obligatoirement créer au minimum les fonctions suivantes :

```go
func afficherMenu()
```

→ affiche le menu.

```go
func obtenirPrix(choix int) int
```

→ retourne le prix de la boisson choisie.

```go
func afficherBoisson(choix int)
```

→ affiche le nom de la boisson choisie.

Le programme doit utiliser une **boucle** afin de permettre à l'utilisateur d'acheter plusieurs boissons.

Si l'utilisateur choisit `0`, le programme s'arrête.

---

### Gestion des erreurs

Si l'utilisateur entre un choix qui n'existe pas :

```text
Choix invalide !
```

Si l'argent inséré est insuffisant :

```text
Montant insuffisant !
Il manque 2 €.
```

Le programme doit alors permettre à l'utilisateur de refaire un choix.

---

### Exemple

```text
=== DISTRIBUTEUR ===

1 - Eau       : 1 €
2 - Soda      : 2 €
3 - Café      : 2 €
4 - Chocolat  : 3 €
0 - Quitter

Votre choix : 4

Vous avez choisi : Chocolat
Prix : 3 €

Montant inséré : 2

Montant insuffisant !
Il manque 1 €.

=== DISTRIBUTEUR ===

Votre choix : 1

Vous avez choisi : Eau
Prix : 1 €

Montant inséré : 5

Merci !
Votre monnaie : 4 €
```
---

## Exercice 2 — Le gestionnaire de notes

**Difficulté : ⭐⭐⭐⭐**  
****Notation : 12pts****

Vous devez créer un programme permettant à un professeur de saisir plusieurs notes et d'obtenir un bilan.

Le programme commence par demander le nombre de notes à saisir.

```text
=== GESTIONNAIRE DE NOTES ===

Combien de notes voulez-vous saisir ? 5
```

Puis il demande chaque note :

```text
Note 1 : 12
Note 2 : 15
Note 3 : 8
Note 4 : 17
Note 5 : 10
```

Le programme doit ensuite afficher :

```text
=== RÉSULTATS ===

Moyenne : 12.4
Note maximale : 17
Note minimale : 8
```

Puis indiquer le résultat :

```text
L'étudiant est admis.
```

Un étudiant est **admis si sa moyenne est supérieure ou égale à 10**.

---

### Contraintes

Vous devez créer au minimum les fonctions suivantes :

```go
func calculerMoyenne(somme int, nombre int) float64
```
Cette fonction retourne la moyenne.

```go
func trouverMaximum(...)
```
Cette fonction permet de trouver la note la plus élevée.

```go
func trouverMinimum(...)
```
Cette fonction permet de trouver la note la plus faible.

```go
func afficherResultat(moyenne float64)
```

Cette fonction affiche :

```text
Admis
```

ou :

```text
Non admis
```

---

### Contraintes supplémentaires

Les notes doivent être comprises entre :

```text
0 et 20
```

Si l'utilisateur saisit une note invalide :

```text
Note invalide ! Une note doit être comprise entre 0 et 20.
```

La saisie doit être redemandée.

---

### Exemple complet

```text
=== GESTIONNAIRE DE NOTES ===

Combien de notes voulez-vous saisir ? 5

Note 1 : 12
Note 2 : 15
Note 3 : 8
Note 4 : 17
Note 5 : 10

=== RÉSULTATS ===

Moyenne : 12.40
Note maximale : 17
Note minimale : 8

Étudiant admis !
```

---

### Bonus (+1)

Ajouter une fonction :

```go
func compterNotesAuDessusDeLaMoyenne(...)
```

qui permet d'afficher le nombre de notes supérieures ou égales à la moyenne.

Exemple :

```text
Nombre de notes au-dessus de la moyenne : 3
```

Vous pouvez également afficher une appréciation :

```text
Moyenne < 10       → Insuffisant
10 à 11.99         → Passable
12 à 13.99         → Assez bien
14 à 15.99         → Bien
16 à 20            → Très bien
```