package main

import "fmt"

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

func (p *Acolyte) GetNom() string {
	return p.Nom
}

func (p *Acolyte) Donne(objet *Objet, destinataire Humain) {
	fmt.Printf("%s donne %s à %s \n", p.Nom, objet.Nom, destinataire.GetNom())
}
