package main

import (
	"reflect"
	"sort"
)

func main() {
}

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	m1 := map[byte]int{}
	m2 := map[byte]int{}

	for i := 0; i < len(word1); i++ {
		m1[word1[i]] += 1
		m2[word2[i]] += 1
	}

	s1 := []int{}
	s2 := []int{}

	sk1 := []int{}
	sk2 := []int{}

	for k, v := range m1 {
		s1 = append(s1, v)
		sk1 = append(sk1, int(k))
	}

	for k, v := range m2 {
		s2 = append(s2, v)
		sk2 = append(sk2, int(k))
	}

	sort.Ints(s1)
	sort.Ints(s2)

	sort.Ints(sk1)
	sort.Ints(sk2)

	if !reflect.DeepEqual(sk1, sk2) {
		return false
	}

	return reflect.DeepEqual(s1, s2)

}
