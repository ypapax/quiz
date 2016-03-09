package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortByLength(t *testing.T) {
	actual := sortByLength([]string{"1", "11"})
	if !reflect.DeepEqual(actual, []string{
		"11", "1",
	}) {
		fmt.Println(actual)
		t.Error("not valid sorting")
	}
}
