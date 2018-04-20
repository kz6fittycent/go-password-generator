package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Read the words to an array
func fileToArray(inputFile *os.File) []string {
	fileReader := bufio.NewReader(inputFile)
	scanner := bufio.NewScanner(fileReader)
	scanner.Split(bufio.ScanLines)
	var fileArray []string
	for scanner.Scan() {
		line := scanner.Text()
		fileArray = append(fileArray, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return fileArray
}

// Pick a pseudorandom number
func randomNumber(totalLines int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	lineNumber := rand.Intn(totalLines)
	return lineNumber
}

// Return word from array
func getWord(fileContents []string) string {
	wordLength := len(fileContents) - 1
	wordRandomNumber := randomNumber(wordLength)
	word := fileContents[wordRandomNumber]
	return word
}

// Transform the words
func transformWord(word string) string {
	wordArray := strings.Split(word, "")
	wordArray[0] = strings.ToUpper(wordArray[0])

	// would this be more efficient using regex?
	for i := range wordArray {
		if i != 0 {
			switch wordArray[i] {
			case "e":
				wordArray[i] = "3"
			case "o":
				wordArray[i] = "0"
			case "a":
				wordArray[i] = "@"
			case "i":
				wordArray[i] = "1"
			case "s":
				wordArray[i] = "$"
			case "g":
				wordArray[i] = "9"
			}
		}
	}

	transformedWord := strings.Join(wordArray, "")
	return transformedWord
}

// Check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Input Files
	noun, err := os.Open("nounlist")
	check(err)
	adjective, err := os.Open("adjectivelist")
	check(err)
	nounArray := fileToArray(noun)
	adjectiveArray := fileToArray(adjective)

	adjectiveWord := getWord(adjectiveArray)
	nounWord := getWord(nounArray)

	fmt.Println(transformWord(adjectiveWord) + transformWord(nounWord))
}
