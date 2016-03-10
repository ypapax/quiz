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
		return
	}
	fileBytes, err := ioutil.ReadFile(wordsListPath)
	if err != nil {
		logErr(err)
		return
	}
	words := strings.Split(string(fileBytes), "\n")
	fmt.Println("inputted words count", len(words))
	compound, parts := FindLongestCompound(words)

	if len(compound) == 0 {
		fmt.Println("no result")
		return
	}
	fmt.Println("The longest compound word is")
	fmt.Println(compound)
	fmt.Println("It has", len(compound), "characters")

	fmt.Println("It consists of the following words:")
	fmt.Println(parts)
}

func FindLongestCompound(words []string) (compound string, parts []string) {
	// sort words by length
	// longest goes first
	sortedWords := sortByLength(words)
	for i, compoundWordCandidate := range sortedWords {
		var subWords []string
		for j, partWordCandidate := range sortedWords {
			if j <= i {
				continue
			}
			if compoundWordCandidate == partWordCandidate {
				continue
			}
			// not interested in empty sub words
			if len(partWordCandidate) == 0 {
				continue
			}
			if !strings.Contains(compoundWordCandidate, partWordCandidate) {
				continue
			}

			if contains(subWords, partWordCandidate) {
				continue
			}

			subWords = append(subWords, partWordCandidate)
			// a word is minimum compound of 2
			if len(subWords) < 2 {
				continue
			}
			validParts := getCompoundParts(compoundWordCandidate, subWords)
			if len(validParts) == 0 {
				continue
			}
			compound = compoundWordCandidate
			parts = validParts
			return
		}
	}
	return
}

func contains(slice []string, newItem string) bool {
	for _, item := range slice {
		if item == newItem {
			return true
		}
	}
	return false
}

func getBeginEndInternal(wholeWord string, parts []string) (begin, end, internal string, internalPartsCandidates []string) {
	if len(parts) == 0 {
		return
	}
	if len(parts) == 1 {
		if wholeWord == parts[0] {
			begin = parts[0]
			end = begin
			return
		}
	}
	var begin_i, end_i int
	for i, part := range parts {
		if strings.HasPrefix(wholeWord, part) {
			begin = part
			begin_i = i
		}
		if strings.HasSuffix(wholeWord, part) {
			end = part
			end_i = i
		}
		if len(begin) > 0 && len(end) > 0 {
			break
		}
	}
	if len(begin) == 0 || len(end) == 0 {
		return
	}

	internal = strings.Replace(wholeWord, begin, "", 1)
	internal = strings.Replace(internal, end, "", 1)
	if end_i < begin_i {
		end_i, begin_i = begin_i, end_i
	}
	parts = removeElement(parts, begin_i)
	if begin != end {
		parts = removeElement(parts, end_i-1)
	}
	internalPartsCandidates = parts
	return
}

func getCompoundParts(wholeWord string, parts []string) (validParts []string) {
	begin, end, internal, internalPartsCandidates := getBeginEndInternal(wholeWord, parts)
	if len(begin) == 0 || len(end) == 0 {
		return
	}
	if begin == end && internal == "" {
		validParts = []string{begin}
		return
	}
	if begin+end == wholeWord {
		validParts = []string{begin, end}
		return
	}

	validInternalWordParts := getCompoundParts(internal, internalPartsCandidates)
	if len(validInternalWordParts) == 0 {
		return
	}
	validParts = append(validParts, begin)
	validParts = append(validParts, validInternalWordParts...)
	validParts = append(validParts, end)
	return
}

func lettersCount(parts []string) (count int) {
	for _, part := range parts {
		count += len(part)
	}
	return
}

func removeElement(slice []string, i int) []string {
	if i < len(slice)-1 {
		slice = append(slice[:i], slice[i+1:]...)
	} else {
		slice = slice[:i]
	}
	return slice
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
