package main

import (
	"fmt"
)

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

func (p *Pretre) Dit(parole string) {
	fmt.Printf("%s dit: %s \n", p.Nom, parole)
}

func (p *Pretre) GetNom() string {
	return p.Nom
}

func (p *Pretre) Consacre(objet *Objet, enType string) *Espece {
	fmt.Printf("L'Esprit Saint transsubstancie %s en %s \n", objet.Nom, enType)
	return objet.Transsubstancie(enType)
}
