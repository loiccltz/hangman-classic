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

func ToFind() {

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	word := Word[rand.Intn(len(Word))]
	lives := 10

	blanks := make([]string, len(word))
		for i := range blanks {
			blanks[i] = "_"
		}
		
		LetterToShow := len(word)/2 - 1
		
		shownIndices := make(map[int]bool)
		
		for i := 0; i < LetterToShow; i++ {
			randIndex := rand.Intn(len(word))
		
			// S'assure que l'indice n'est pas déjà utilisé
			for shownIndices[randIndex] {
				randIndex = rand.Intn(len(word))
			}
			
			blanks[randIndex] = string(word[randIndex]) // Révèle la lettre
			shownIndices[randIndex] = true              // Marque cet indice comme révélé
		}
	


	for {
		// 3. met les "_"
		fmt.Printf("❤️ %d, Word: %s Letter: ", lives, strings.Join(blanks, " "))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// 4. verifie si la lettre est bonne
		for _, inputLetter := range input {
			correctGuess := false
			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					blanks[i] = string(inputLetter)
					correctGuess = true
				}
			}

			if !correctGuess {
				lives--
				
			}

		}

		// 5. plus de vie ff
		if lives <= 0 {
			fmt.Printf("❤️ 0, Word: %s - FF!\n", word)
			break
		}
		// 6. si le mot est deviné
		if word == strings.Join(blanks, "") {
			fmt.Printf("❤️ %d, Word: %s - GG!\n", lives, word)
			break
		}
	}
}

