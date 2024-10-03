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
	const (
		heigt = 7
	)
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
	word := Word[rand.Intn(len(Word))] // on recupere les lettre que l'on veut cacher au hasard                     // initialisation des vies
	lives := 10
	blanks := make([]string, len(word))
	for i := range blanks {
		blanks[i] = "_" // on cache les lettres
	}

	LetterToShow := len(word)/2 - 1

	shownIndices := make(map[int]bool)

	for i := 0; i < LetterToShow; i++ {
		randIndex := rand.Intn(len(word))

		for shownIndices[randIndex] {
			randIndex = rand.Intn(len(word))
		}

		blanks[randIndex] = string(word[randIndex])
		shownIndices[randIndex] = true
	}

	for {
		// 3. met les "_"
		fmt.Printf("❤️ %d, Word: %s Letter: ", lives, strings.Join(blanks, " "))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// 4. verifie si la lettre est bonne

		file, err := os.Open("positions/hangman.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string
	  
		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
	  
		file.Close()

		lineCount := 0
		for _, line := range fileLines {
			fmt.Println(line)
				lineCount++
			}

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
