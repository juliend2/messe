package main

import (
	"fmt"
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

// Prêtre
type Pretre struct {
	Humain
	Nom      string
	Position string
}

func (p *Pretre) AllerA(position string) {
	p.Position = position
	fmt.Printf("%s (le prêtre) s'en va à %s \n", p.Nom, p.Position)
}

func (p *Pretre) Dire(parole string) {
	fmt.Printf("%s dit: %s \n", p.Nom, parole)
}

func (p *Pretre) GetNom() string {
	return p.Nom
}

func (p *Pretre) Consacre(objet *Objet, enType string) *Espece {
	fmt.Printf("L'Esprit Saint transsubstancie %s en %s \n", objet.Nom, enType)
	return objet.Transsubstancie(enType)
}

// Acolyte

type Acolyte struct {
	Humain
	Nom      string
	Position string
}

func (p *Acolyte) AllerA(position string) {
	p.Position = position
	fmt.Printf("%s (un acolyte) s'en va à %s \n", p.Nom, p.Position)
}

func (p *Acolyte) Lit(passage string) {
	fmt.Printf("%s lit: %s \n", p.Nom, passage)
}

func (p *Acolyte) Donne(objet *Objet, destinataire Humain) {
	fmt.Printf("%s donne %s à %s \n", p.Nom, objet.Nom, destinataire.GetNom())
}

func RitesInitiaux(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	ministres := make([]Humain, 0, 3)
	ministres = append(ministres, celebrant, acolyte1, acolyte2)

	for _, ministre := range ministres {
		ministre.AllerA("l'autel")
	}

	celebrant.Dire("Prions")
}

func LiturgieDeLaParole(celebrant *Pretre, lecteur Humain) {
	lecteur.Lit("Première Épitre aux Corinthiens, 1.5")
}

func LiturgieDeLeucharistie(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	celebrant.AllerA("l'autel")
	acolyte1.AllerA("l'autel")
	acolyte2.AllerA("l'autel")

	pain := &Objet{Nom: "le pain"}
	vin := &Objet{Nom: "le vin"}

	acolyte1.Donne(&Objet{Nom: "le calice"}, celebrant)
	acolyte2.Donne(&Objet{Nom: "le ciboire"}, celebrant)

	acolyte2.Donne(pain, celebrant)
	acolyte2.Donne(vin, celebrant)

	celebrant.Consacre(vin, Sang)
	celebrant.Consacre(pain, Corps)
}

func main() {
	celebrant := &Pretre{Nom: "Claude"}
	acolyte1 := &Acolyte{Nom: "Julien"}
	acolyte2 := &Acolyte{Nom: "Pierre"}

	RitesInitiaux(celebrant, acolyte1, acolyte2)

	LiturgieDeLaParole(celebrant, acolyte1)
	LiturgieDeLeucharistie(celebrant, acolyte1, acolyte2)
}
