package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// voir strings.Contains() pour le mot https://www.geeksforgeeks.org/string-contains-function-in-golang-with-examples/

func Word() {

	file, err := os.Open("dictionnaries/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// je defini word ( en dehors de la boucle)
	var Word []string
	for scanner.Scan() {
		var tab = scanner.Text()
		Word = append(Word, tab) // j'ajoute le contenue de tab a Word
	}

	word := Word[rand.Intn(len(Word))]
	wordRunes := []rune(word)
	lives := 10

	// permet de cacher le mot "php" -> p_p
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_' 
	}

	var allLetters string
	usedLetters := make(map[rune]bool)

	// Variable pour garder une trace du nombre de lignes deja affichées dans le fichier du pendu hangman.txt
	var linesDisplayed int

	for {
		// show the word blanks and ask for letters
		fmt.Printf("\n %d ❤️, Mot: %s, \n", lives, strings.Join(convertRuneSliceToStringSlice(blanks), " ")) // on convertie le slice de rune en slice de string
		fmt.Println(" Mot de : ", len(word), " lettres ", "\n")
		fmt.Println("Lettre déjà proposée : ", allLetters, "\n")

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)


	
	if input == "stop" { // si le joueur entre STOP
		saveData := map[string]interface{}{ // on enregistre l'état de la partie
			"word":           word,
			"lives":          lives,
			"blanks":         blanks,
			"guessedLetters": allLetters,
			"usedLetters":    usedLetters,
		}
		
		file, err := os.Create("save.txt") // on créer le fichier save.txt / met a jour
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		
		jsonData, err := json.Marshal(saveData) // on transforme les données en json
		if err != nil {
			log.Fatal(err)
		}
		file.Write(jsonData)
		break
	}
	if len(input) >= 2 {
		if input == string(wordRunes) {
				// Cas où le joueur entre un mot complet
				// Le joueur a deviné le mot correctement
				fmt.Printf("\n %d ❤️, Mot: %s - Vous avez gagné!\n", lives, string(wordRunes))
				break
			} else {
				// Le mot est incorrect, on retire 2 vies
				lives -= 2
				fmt.Printf("Mot incorrect ! Vous perdez 2 vies.\n")
			}
	} else {

		
		allLetters += input
		
		inputLetter := rune(input[0])
		// check provided letters
		
		if usedLetters[inputLetter] {
			fmt.Printf("Vous avez déjà utilisé cette lettre: %c\n", inputLetter)
			continue
		}
			
			correctGuess := false
			
			for i, wordLetter := range wordRunes {
				if inputLetter == wordLetter || (inputLetter == 'e' && wordLetter == 'é') || (inputLetter == 'é' && wordLetter == 'e') {
					blanks[i] = wordLetter
					correctGuess = true
				}
			}
			
			if !correctGuess {
				lives--
				
				// afficher hangman
				linesDisplayed = showHangman(linesDisplayed)
			}
	}

		// if no more lives, you lost
		if lives <= 0 {
			fmt.Printf("\n 0 ❤️, Mot: %s - Vous avez perdu!\n", string(wordRunes))
			break
		}
		// if word is guessed, you won
		if string(wordRunes) == string(blanks) {
			fmt.Printf("\n %d ❤️, Mot: %s - Vous avez gagné!\n", lives, string(wordRunes))
			break
		}
	}
}

// Fonction pour convertir un slice de rune en slice de string
func convertRuneSliceToStringSlice(runes []rune) []string {
	strings := make([]string, len(runes))
	for i, r := range runes {
		strings[i] = string(r)
	}
	return strings
}

// fonction pour afficher les 8 premieres du pendu
func showHangman(linesDisplayed int) int {
	hangmanFile, err := os.Open("dictionnaries/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer hangmanFile.Close()

	scanner := bufio.NewScanner(hangmanFile)
	lineCount := 0
	startLine := linesDisplayed // Lignes déja affichées

	fmt.Println("\n--- Hangman Status ---")
	for scanner.Scan() {
		// Si on atteint 8 nouvelles lignes on arrete
		if lineCount >= startLine+8 {
			break
		}

		// On saute les lignes déja affichées
		if lineCount < startLine {
			lineCount++
			continue
		}

		// Afficher les 8 prochaines lignes
		if lineCount >= startLine && lineCount < startLine+8 {
			fmt.Println(scanner.Text())
			lineCount++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------------------\n")
	return lineCount // Retourne le nombre total de lignes affichées
}
