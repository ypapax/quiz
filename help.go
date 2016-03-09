package main

import (
	"encoding/json"
	"fmt"
)

func FromJson(jsonSrc string, objRef interface{}) {
	if len(jsonSrc) == 0 {
		return
	}
	err := json.Unmarshal([]byte(jsonSrc), &objRef)
	if err != nil {
		panic(err)
	}
}

func ToJson(obj interface{}) string {
	b, err := json.MarshalIndent(&obj, "", "   ")
	if err != nil {
		panic(err)
	}
	strJson := string(b)
	return strJson
}

func ActualExpected(actual, expected string) {
	fmt.Println("actual")
	fmt.Println(actual)
	fmt.Println("expected")
	fmt.Println(expected)

}
