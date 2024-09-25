package main

import (
	"fmt"
	"reflect"
)

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func getKeys2(m interface{}) []interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("not a map")
	}
	keys := make([]interface{}, 0, v.Len())
	for _, key := range v.MapKeys() {
		keys = append(keys, key.Interface())
	}
	return keys
}

func getKeys3[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	myMap := map[string]interface{}{
		"name": "Alice",
		"age":  38,
		"city": "New York",
	}
	keys := getKeys(myMap)
	fmt.Println(keys)

	myMap2 := map[interface{}]interface{}{
		"name": "Alice",
		123:    40,
		true:   "New York",
	}
	keys2 := getKeys2(myMap2)
	fmt.Println(keys2)

	stringMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   35,
	}
	keys1 := getKeys3(stringMap)
	fmt.Println(keys1)

	intMap := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	keys3 := getKeys3(intMap)
	fmt.Println(keys3)

	mixedMap := map[interface{}]int{
		"Alice": 25,
		123:     30,
		true:    35,
	}
	keys4 := getKeys3(mixedMap)
	fmt.Println(keys4)
}
