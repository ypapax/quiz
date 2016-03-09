package main

import "testing"

func TestFindLongestCompound(t *testing.T) {
	compound, parts := FindLongestCompound([]string{
		"hello",
		"seven",
		"world",
		"one",
		"two",
		"helloworld",
	})
	if compound != "helloworld" {
		t.Error("not valid compound: ", compound)
		return
	}
	if len(parts) != 2 {
		t.Error("not valid parts")
		return
	}

	if parts[0] != "hello" {
		t.Error("not valid first word")
	}
	if parts[1] != "world" {
		t.Error("not valid second word")
	}
}

func Test_isCompound(t *testing.T) {
	validPermutation := isCompound("bedroomman", []string{"man", "bed", "room"})
	expectedJson := `[
   "bed",
   "room",
   "man"
]`
	if actualJson := ToJson(validPermutation); actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("isCompound not working")
	}
}
