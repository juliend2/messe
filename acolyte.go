package main

import "fmt"

// Acolyte

type Acolyte struct {
	Humain
	Nom      string
	Position string
}

func (a *Acolyte) AllerA(position string) {
	a.Position = position
	fmt.Printf("%s (un acolyte) s'en va à %s \n", a.Nom, a.Position)
}

func (a *Acolyte) Lit(passage string) {
	fmt.Printf("%s lit: %s \n", a.Nom, passage)
}

func (a *Acolyte) GetNom() string {
	return a.Nom
}

func (a *Acolyte) Donne(objet *Objet, destinataire Humain) {
	fmt.Printf("%s donne %s à %s \n", a.Nom, objet.Nom, destinataire.GetNom())
}

func (a *Acolyte) MetSurAutel(objet Objets) {
	// if typeof(objet) == "Objet" {
	// 	objetString := objet.Nom
	// } else {
	// 	objetString := JoinAvecEt(Map(objets, func (obj) string { return obj.Nom }))
	// }
	// fmt.Printf("%s met sur l'autel ", objetString)
	fmt.Printf("%s met sur l'autel les %s objets \n", a.Nom, objet.Nom)
}
