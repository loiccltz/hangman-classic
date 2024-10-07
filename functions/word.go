package hangman

import (
	"bufio"
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
	lives := 10

	//generate the word blanks  "php" -> p_p
	blanks := []string{}
	for range word {
		blanks = append(blanks, "_")
	}

	var allLetters string
	usedLetters := make(map[rune]bool)

	// Variable pour garder une trace du nombre de lignes affichées
	var linesDisplayed int

	for {
		// show the word blanks and ask for letters
		fmt.Printf("\n %d ❤️, Word: %s, \n", lives, strings.Join(blanks, " "))
		fmt.Print(len(word), " Word letter: ", allLetters)

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		allLetters += input

		// check provided letters
		for _, inputLetter := range input {
			if usedLetters[inputLetter] {
				fmt.Printf("You have already used the letter: %c\n", inputLetter)
				continue
			}

			usedLetters[inputLetter] = true

			correctGuess := false

			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					blanks[i] = string(inputLetter)
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
			fmt.Printf("\n 0 ❤️, Word: %s - sorry, you lost!\n", word)
			break
		}
		// if word is guessed, you won
		if word == strings.Join(blanks, "") {
			fmt.Printf("\n %d ❤️, Word: %s - you won, congrats!\n", lives, word)
			break
		}
	}
}

// Fonction pour afficher les 8 premières lignes de hangman.txt
func showHangman(linesDisplayed int) int {
	hangmanFile, err := os.Open("dictionnaries/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer hangmanFile.Close()

	scanner := bufio.NewScanner(hangmanFile)
	lineCount := 0
	startLine := linesDisplayed // Lignes déjà affichées

	fmt.Println("\n--- Hangman Status ---")
	for scanner.Scan() {
		// Si on atteint 8 nouvelles lignes, on arrête
		if lineCount >= startLine+8 {
			break
		}

		// On saute les lignes déjà affichées
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
