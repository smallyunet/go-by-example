package main

import "fmt"

func getKeys1(m map[string]interface{}) []string {
	k := make([]string, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	return k
}

func getKeys2(m map[interface{}]interface{}) []interface{} {
	k := make([]interface{}, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	return k
}

func getKeys3[K comparable, V any](m map[K]V) []K {
	k := make([]K, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	return k
}

func main() {
	m1 := map[string]interface{}{
		"a": "a",
		"b": 1,
		"c": true,
	}
	m2 := map[interface{}]interface{}{
		"a":  1,
		1:    2,
		true: 3,
	}
	m3 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	k1 := getKeys1(m1)
	k2 := getKeys2(m2)
	k3 := getKeys3(m3)
	fmt.Println(k1, k2, k3)
}
