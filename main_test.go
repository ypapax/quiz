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

func Test_getCompoundParts_(t *testing.T) {
	validParts := getCompoundParts("bedroomman", []string{"man", "bed", "room"})
	expectedJson := `[
   "bed",
   "room",
   "man"
]`
	if actualJson := ToJson(validParts); actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("getCompoundParts not working")
	}
}

func Test_getCompoundParts4(t *testing.T) {
	validParts := getCompoundParts("bedroomman", []string{"man", "dro", "bed", "room"})
	expectedJson := `[
   "bed",
   "room",
   "man"
]`
	if actualJson := ToJson(validParts); actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("getCompoundParts not working")
	}
}

func Test_getCompoundParts12(t *testing.T) {
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
	validParts := getCompoundParts("pneumonoultramicroscopicsilicovolcanoconiosis", input)

	if len(validParts) > 0 {
		t.Error("must be empty")
	}
}

func Test_getCompoundParts31(t *testing.T) {
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
	validParts := getCompoundParts("dichlorodiphenyltrichloroethanes", input)

	if len(validParts) > 0 {
		t.Error("must be empty")
	}
}

func Test_getCompoundParts29(t *testing.T) {
	inputJson := `[
   "dis",
   "anti",
   "establishmentarianisms"
]`
	var input []string
	FromJson(inputJson, &input)
	validParts := getCompoundParts("antidisestablishmentarianisms", input)
	expectedJson := `[
   "anti",
   "dis",
   "establishmentarianisms"
]`
	if actualJson := ToJson(validParts); actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("not valid")
	}
}

func Test_getCompoundParts29Full(t *testing.T) {
	inputJson := `[
   "establishmentarianisms",
   "establishmentarianism",
   "establishment",
   "stablishment",
   "establish",
   "stablish",
   "menta",
   "ment",
   "stab",
   "anis",
   "anti",
   "aria",
   "ism",
   "ish",
   "nis",
   "lis",
   "est",
   "tar",
   "tab",
   "tid",
   "dis",
   "ani",
   "ria",
   "ant",
   "men",
   "dis",
   "sh",
   "en",
   "id",
   "hm",
   "es",
   "st",
   "ab",
   "me",
   "ta",
   "li",
   "an",
   "is",
   "ti",
   "di",
   "ar"
]`
	var input []string
	FromJson(inputJson, &input)
	validParts := getCompoundParts("antidisestablishmentarianisms", input)
	expectedJson := `[
   "anti",
   "dis",
   "establishmentarianisms"
]`
	if actualJson := ToJson(validParts); actualJson != expectedJson {
		ActualExpected(actualJson, expectedJson)
		t.Error("not valid")
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

func Test_getBeginEndInternal29(t *testing.T) {
	wholeWord := "antidisestablishmentarianisms"
	inputJson := `[
   "establishmentarianisms",
   "establishmentarianism",
   "establishment",
   "stablishment",
   "establish",
   "stablish",
   "menta",
   "ment",
   "stab",
   "anis",
   "anti",
   "aria",
   "ism",
   "ish",
   "dis",
   "nis",
   "lis",
   "est",
   "tar",
   "tab",
   "tid",
   "ani",
   "ria",
   "ant",
   "men",
   "dis",
   "sh",
   "en",
   "id",
   "hm",
   "es",
   "st",
   "ab",
   "me",
   "ta",
   "li",
   "an",
   "is",
   "ti",
   "di",
   "ar"
]`
	var input []string
	FromJson(inputJson, &input)

	begin, end, internal, _ := getBeginEndInternal(wholeWord, input)
	if begin != "anti" {
		t.Error("not valid begin", begin)
	}
	if end != "establishmentarianisms" {
		t.Error("not valid end")
	}
	if internal != "dis" {
		t.Error("not valid internal", internal)
	}
}

func Test_contains(t *testing.T) {
	inputJson := `[
   "establishmentarianism",
   "establishmentarianism",
   "establishmentarianism",
   "ar"
]`
	var input []string
	FromJson(inputJson, &input)
	if !contains(input, "establishmentarianism") {
		t.Error("contains not working")
	}
}
