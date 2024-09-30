package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
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
		var tab  =  scanner.Text()
		Word = append(Word, tab) // j'ajoute le contenue de tab a Word
    }
	RandomWord := rand.Intn(len(Word)) // je genere un rand longueur max de 37 (nombre total de mots)
	fmt.Println(Word[RandomWord]) 
	

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
}
