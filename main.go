package main

import (
	"fmt"
	"strings"
)

const (
	Sang  = "sang du Christ"
	Corps = "corps du Christ"
)

type Espece struct {
	Nom  string
	Type string
}

type Humain interface {
	GetNom() string
	AllerA(position string)
	Dire(parole string)
	Lit(passage string)
	Donne(objet *Objet, destinataire Humain)
}

// Objet

type Objet struct {
	Nom string
}

func (o *Objet) Transsubstancie(typeDeDestination string) *Espece {
	return &Espece{Nom: o.Nom, Type: typeDeDestination}
}

func RitesInitiaux(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	ministres := make([]Humain, 0, 3)
	ministres = append(ministres, celebrant, acolyte1, acolyte2)
	MarcheDesMinistres(ministres, "l'autel")

	celebrant.Dire("Prions")
}

func LiturgieDeLaParole(celebrant *Pretre, lecteur Humain) {
	lecteur.Lit("Première Épitre aux Corinthiens, 1.5")
}

func LiturgieDeLeucharistie(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	MarcheDesMinistres([]Humain{acolyte1, acolyte2, celebrant}, "l'autel")

	pain := &Objet{Nom: "le pain"}
	vin := &Objet{Nom: "le vin"}

	acolyte1.Donne(&Objet{Nom: "le calice"}, celebrant)
	acolyte2.Donne(&Objet{Nom: "le ciboire"}, celebrant)

	acolyte2.Donne(pain, celebrant)
	acolyte2.Donne(vin, celebrant)

	celebrant.Consacre(vin, Sang)
	celebrant.Consacre(pain, Corps)
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

func Map[V any, R any](input []V, fn func(V) R) []R {
	results := make([]R, len(input))
	for i, item := range input {
		results[i] = fn(item)
	}
	return results
}

func MarcheDesMinistres(ministres []Humain, versPosition string) {
	noms := JoinAvecEt(Map(ministres, func(m Humain) string { return m.GetNom() }))
	fmt.Printf("%s se dirigent vers %s \n", noms, versPosition)
}

func main() {
	celebrant := &Pretre{Nom: "Claude"}
	acolyte1 := &Acolyte{Nom: "Julien"}
	acolyte2 := &Acolyte{Nom: "Pierre"}

	RitesInitiaux(celebrant, acolyte1, acolyte2)

	LiturgieDeLaParole(celebrant, acolyte1)
	LiturgieDeLeucharistie(celebrant, acolyte1, acolyte2)
}
