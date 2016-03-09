package main

import (
	"fmt"
	"testing"
)

func Test_getPermutations(t *testing.T) {
	actual := getPermutations([]string{"1", "2", "3"})
	if len(actual) != 6 {
		t.Error("not valid permutations count")
		fmt.Println(actual)
	}
	expectedJson := `[
   [
      "1",
      "2",
      "3"
   ],
   [
      "1",
      "3",
      "2"
   ],
   [
      "2",
      "1",
      "3"
   ],
   [
      "2",
      "3",
      "1"
   ],
   [
      "3",
      "1",
      "2"
   ],
   [
      "3",
      "2",
      "1"
   ]
]`
	actualJson := ToJson(actual)
	if actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("not valid permutations")
	}

}
