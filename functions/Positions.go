package hangman

import (
	"log"
	"os"
	"bufio"
	"fmt"
)
	
	const (
		attempts = 10 // Nombre de tentatives
		height   = 7  // Hauteur de chaque position
		width	 = 9
	)
		

	
	func Positions() {
		// Ouvrir le fichier hangman.txt
		
		file, err := os.Open("/home/alexandre-reffert/Bureau/Projet-Hangman/hangman-classic/positions/hangman.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		// Lire les positions
		var positions [attempts][height]string
		scanner := bufio.NewScanner(file)
	
		for i := 0; i < attempts; i++ {
			for j := 0; j < height; j++ {
				if scanner.Scan() {
					positions[i][j] = scanner.Text()
				}
			}
			// Lire la ligne vide qui suit chaque position
			scanner.Scan()
		}

		
		errorCount := 0
	
		// Simuler le jeu
		for remainingAttempts := attempts; remainingAttempts > 0; remainingAttempts-- {
			var guess string
			fmt.Print("Entrez une lettre : ")
		fmt.Scan(&guess)

		if guess != "correctLetter" { // Remplacez "correctLetter" par votre lettre correcte
			errorCount++
			if errorCount <= attempts {
				// Afficher la position de José en fonction des erreurs
				JosePosition(positions[errorCount-1])
			}
			fmt.Printf("Mauvaise réponse ! Tentatives restantes : %d\n", remainingAttempts-1)
		} else {
			fmt.Println("Bonne réponse !")
		}
	}
}
	
func JosePosition(position [height]string) {
	for _, line := range position {
			fmt.Println(line)
	}
}
