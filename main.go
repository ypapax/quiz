package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	wordsListPath := getArgument(1)
	if len(wordsListPath) == 0 {
		fmt.Println("usage:")
		fmt.Println("quiz /path/to/word.list")
	}
	fileBytes, err := ioutil.ReadFile(wordsListPath)
	if err != nil {
		logErr(err)
		return
	}
	words := strings.Split(string(fileBytes), "\n")
	compound, parts := FindLongestCompound(words)
	if len(compound) == 0 {
		fmt.Println("no result")
		return
	}
	fmt.Println("The longest compound word is")
	fmt.Println(compound)
	fmt.Println("It consists of the following words:")
	fmt.Println(parts)
}

func FindLongestCompound(words []string) (compound string, parts []string) {
	// sort words by length
	// longest goes first
	sortedWords := sortByLength(words)
	resultCandidates := make(map[string][]string)
	for _, compoundWordCanidate := range sortedWords {
		for _, partWordCandidate := range sortedWords {
			if compoundWordCanidate == partWordCandidate {
				continue
			}
			if !strings.Contains(compoundWordCanidate, partWordCandidate) {
				continue
			}
			resultCandidates[compoundWordCanidate] = append(
				resultCandidates[compoundWordCanidate], partWordCandidate)
			validPermutation := isCompound(compoundWordCanidate, resultCandidates[compoundWordCanidate])
			if len(validPermutation) == 0 {
				continue
			}
			compound = compoundWordCanidate
			parts = validPermutation
		}
	}
	return
}

func isCompound(wholeWord string, parts []string) (validPermutation []string) {
	if len(parts) < 2 {
		return
	}
	permutations := getPermutations(parts)

	for _, permutation := range permutations {
		if wholeWord == strings.Join(permutation, "") {
			validPermutation = permutation
			return
		}

	}
	return
}

func getArgument(argNumber int) (argValue string) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < argNumber {
		return
	}
	argValue = argsWithoutProg[argNumber-1]
	return
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
