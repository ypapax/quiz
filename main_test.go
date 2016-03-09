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

func Test_isCompound_(t *testing.T) {
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

func Test_isCompound4(t *testing.T) {
	validPermutation := isCompound("bedroomman", []string{"man", "dro", "bed", "room"})
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

func Test_isCompound12(t *testing.T) {
	inputJson := `[
   "ultramicroscopic",
   "microscopic",
   "ultramicro",
   "coniosis",
   "volcano",
   "micros",
   "ultra",
   "micro",
   "mono",
   "scop",
   "tram",
   "rami"
]`
	var input []string
	FromJson(inputJson, &input)
	validPermutation := isCompound("pneumonoultramicroscopicsilicovolcanoconiosis", input)

	if len(validPermutation) > 0 {
		t.Error("must be empty")
	}
}

func Test_isCompound31(t *testing.T) {
	inputJson := `[
   "diphenyl",
   "ethanes",
   "ethane",
   "phenyl",
   "thane",
   "thane",
   "rich",
   "rich",
   "rich",
   "than",
   "anes",
   "lor",
   "eth",
   "ane",
   "dip",
   "han",
   "ich",
   "roe",
   "rod",
   "hen",
   "ny",
   "en",
   "od",
   "ch",
   "ph",
   "lo",
   "et",
   "an",
   "ne",
   "he",
   "ha"
]`
	var input []string
	FromJson(inputJson, &input)
	validPermutation := isCompound("dichlorodiphenyltrichloroethanes", input)

	if len(validPermutation) > 0 {
		t.Error("must be empty")
	}
}

func Test_getBeginEndInternal(t *testing.T) {
	begin, end, internal, internalPartsCandidates := getBeginEndInternal("onetwo", []string{"one", "two"})
	if begin != "one" {
		t.Error("not valid begin")
	}
	if end != "two" {
		t.Error("not valid end")
	}
	if internal != "" {
		t.Error("not valid internal")
	}
	if len(internalPartsCandidates) != 0 {
		t.Error("not valid internalPartsCandidates", internalPartsCandidates)
	}
}

func Test_getBeginEndInternalOne(t *testing.T) {
	begin, end, internal, internalPartsCandidates := getBeginEndInternal("one", []string{"one", "two"})
	if begin != "one" {
		t.Error("not valid begin")
	}
	if end != "one" {
		t.Error("not valid end", end)
	}
	if internal != "" {
		t.Error("not valid internal")
	}
	if len(internalPartsCandidates) != 1 {
		t.Error("not valid internalPartsCandidates", internalPartsCandidates)
	}
}

func Test_getBeginEndInternalDirect(t *testing.T) {
	begin, end, internal, internalPartsCandidates := getBeginEndInternal("one", []string{"one"})
	if begin != "one" {
		t.Error("not valid begin", begin)
	}
	if end != "one" {
		t.Error("not valid end", end)
	}
	if internal != "" {
		t.Error("not valid internal")
	}
	if len(internalPartsCandidates) != 0 {
		t.Error("not valid internalPartsCandidates", internalPartsCandidates)
	}
}
