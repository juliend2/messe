package main

import (
	"fmt"
	"strings"
)

func Map[V any, R any](input []V, fn func(V) R) []R {
	results := make([]R, len(input))
	for i, item := range input {
		results[i] = fn(item)
	}
	return results
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func JoinAvecEt(elements []string) string {
	if len(elements) == 0 {
		return ""
	}
	if len(elements) == 1 {
		return elements[0]
	}

	return strings.Join(elements[:len(elements)-1], ", ") + " et " + elements[len(elements)-1]
}

func MarcheDesMinistres(ministres []Humain, versPosition string) {
	noms := JoinAvecEt(Map(ministres, func(m Humain) string { return m.GetNom() }))
	fmt.Printf("%s se dirigent vers %s \n", noms, versPosition)
}
