# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

# Go Solution

## Build
`go test && go build`

## Run
```
./quiz /path/to/word.list
inputted words count 263534
The longest compound word is
electroencephalographically
It has 27 characters
It consists of the following words:
[electro encephalographically]
```

## Algorithm
1. Since we want the biggest compound word, sort all the words by length - biggest goes first
2. For each word starting from the first (biggest), find other words which first word contains and put this sub words to `subWords` slice.
```
  for _, compoundWordCandidate := range sortedWords {
    var subWords []string
    for _, partWordCandidate := range sortedWords {
      if compoundWordCandidate == partWordCandidate {
        continue
      }
      if !strings.Contains(compoundWordCandidate, partWordCandidate) {
        continue
      }
      subWords = append(subWords, partWordCandidate)
      validParts := getCompoundParts(compoundWordCandidate, subWords)
      if len(validParts) == 0 {
        continue
      }
      compound = compoundWordCandidate
      parts = validParts
      return
    }
  }
```
3. After every append to `subWords`, check if the slice has length more than 1 and try to build the current `compoundWordCandidate` from `subWords` using method `getCompoundParts`. If it returns positive length slice, that's it. Current `compoundWordCandidate` is the result of the program - the longest compound word. It consists of `validParts`.

### `getCompoundParts` algorithm
Function `func getCompoundParts(wholeWord string, parts []string) (validParts []string)` cuts off beginning and end of `wholeWord` reducing number of possible `parts` by 2 and calls itself with shorter `wholeWord` and lower length `parts` until has trivial cases when `parts` length is one or two.

# Solution author
Maxim Yefremov [upwork profile](https://www.upwork.com/o/profiles/users/_~012ca70e652c74ed7c/)