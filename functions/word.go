package hangman

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)


func Word() {

	SaveLoad := flag.Bool("startWidth", false, "Permet de charger un fichier de sauvegarde")
	flag.Parse()

	// je lit le document word.txt
	file, err := os.Open("dictionnaries/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// je defini word (en dehors de la boucle)
	var Word []string
	for scanner.Scan() {
		var tab = scanner.Text()
		Word = append(Word, tab) // j'ajoute le contenue de tab a Word
	}

	word := Word[rand.Intn(len(Word))] // je prend le nombre de lettre dans le mot
	wordRunes := []rune(word)
	lives := 10 // je met la vie a 10

	// permet de cacher chaque lettres du mot par "_"
	blanks := make([]rune, len(wordRunes))
	for i := range blanks {
		blanks[i] = '_'
	}

	// je laisse imprimé la/les lettres mis en input
	var allLetters string
	usedLetters := make(map[rune]bool)

	// Variable pour garder une trace du nombre de lignes deja affichées dans le fichier du pendu hangman.txt
	var linesDisplayed int

	if *SaveLoad { // si le flag est utilisé
		file, err := os.Open("save.txt") // on charge la sauvegarde
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var saveData map[string]interface{} // interface permet de stocker tout type de data

		json.NewDecoder(file).Decode(&saveData)

		word = saveData["word"].(string) // je recupere la valeur word, en m'assurant qu'elle soit de type string

		lives = int(saveData["lives"].(float64)) // on transforme en int

		blanks = []rune(saveData["blanks"].(string)) // on transforme de string en rune

		// on reforme les lettres utilisé
		for _, letter := range saveData["usedLetters"].(string) {
			usedLetters[rune(letter)] = true
		}

		allLetters = saveData["guessedLetters"].(string)
	}

	// Correspondance des lettres avec accents
	SpecialLetters := map[rune][]rune{
		'a': {'a', 'à', 'â'},
		'e': {'e', 'é', 'ê', 'è', 'ë'},
		'i': {'i', 'ï', 'î'},
		'o': {'o', 'ô'},
		'u': {'u', 'ù', 'û'},
	}

	for {
		// j'imprime le mot en blank, le nombre de vie, le nombre de lettres du mot et les lettres déja utiliser
		fmt.Printf("\n %d ❤️, Word : %s, \n", lives, strings.Join(convertRuneSliceToStringSlice(blanks), " "))
		fmt.Print("Word of : ", len([]rune(word)), " letters: ", "\n")
		fmt.Print("You have already used the letters : ", allLetters, "\n")

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// si le joueur entre STOP le programme s'arrete
		if input == "stop" { // si le joueur entre STOP
			saveData := map[string]interface{}{ // on enregistre l'état de la partie

				// clé ici string	// type donc ici interface
				"word":           word,
				"lives":          lives,
				"blanks":         string(blanks), // on stock en string sinon cela ne marche pas avec le JSON
				"guessedLetters": allLetters,
				"usedLetters":    string(getKeysFromMap(usedLetters)),
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
			fmt.Println("Votre partie a bien été sauvegardé")
			break
		}

		// Permet d'entrer un mot de 2 lettres ou plus
		if len(input) > 2 {
			if removeAccents(input) == removeAccents(word) {
				fmt.Printf("\n %d ❤️, Word: %s - You won, congrats!\n", lives, word)
				break
			} else {
				lives -= 2
				fmt.Printf("\nWrong word! You lost 2 lives. %d ❤️ remaining.\n", lives)
				linesDisplayed = showHangman(linesDisplayed)
				linesDisplayed = showHangman(linesDisplayed)

				if lives <= 0 {
					fmt.Printf("\n 0 ❤️, Word: %s - Sorry, you lost!\n", word)
					break
				}
			}
		} else {

			allLetters += input

			for _, inputLetter := range input {
				if usedLetters[inputLetter] {
					fmt.Printf("\nYou have already used the letter: %c\n", inputLetter)
					continue
				}

				usedLetters[inputLetter] = true

				correctGuess := false

				for i, wordLetter := range wordRunes {
					// on verifie les accents si ils existent dans specialLetters
					if letters, exists := SpecialLetters[inputLetter]; exists {
						for _, possibleLetter := range letters {
							if wordLetter == possibleLetter {
								blanks[i] = wordLetter
								correctGuess = true
							}
						}
					} else {
						// Sinon, on vérifie la lettre directement
						if wordLetter == inputLetter {
							blanks[i] = wordLetter
							correctGuess = true
						}
					}
				}

				// Si la lettre n'est pas dans le mot on enlève une vie
				if !correctGuess {
					lives--
					fmt.Printf("\nWrong guess! %d ❤️ remaining.\n", lives)
					linesDisplayed = showHangman(linesDisplayed)
				}
			}

			// si vie = 0 alors on perd
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
}

func getKeysFromMap(m map[rune]bool) []rune {
	keys := []rune{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func convertRuneSliceToStringSlice(runes []rune) []string {
	strings := make([]string, len(runes))
	for i, r := range runes {
		strings[i] = string(r)
	}
	return strings
}

func removeAccents(input string) string {
	replacer := strings.NewReplacer(
		"à", "a", "â", "a",
		"é", "e", "ê", "e", "è", "e", "ë", "e",
		"ï", "i", "î", "i",
		"ô", "o",
		"ù", "u", "û", "u",
	)
	return replacer.Replace(input)
}

// Fonction pour afficher le hangman
func showHangman(linesDisplayed int) int {
	hangmanFile, err := os.Open("dictionnaries/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer hangmanFile.Close()

	scanner := bufio.NewScanner(hangmanFile)

	lineCount := 0 // on commence à la ligne 0 du fichier txt

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

	fmt.Printf("----------------------\n")
	return lineCount
}
