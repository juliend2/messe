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

	Dit(parole string)
	AllerA(position string)
	Lit(passage string)
	Donne(objet *Objet, destinataire Humain)
	MetSurAutel(objet Objets) // Objet ou Objets
}

type Peuple struct {
	Humain
	Nom string
}

func (p *Peuple) Dit(parole string) {
	fmt.Printf("%s dit: %s \n", p.Nom, parole)
}

// Objet

type Objet struct {
	Nom string
}

type Objets struct {
	Nom string
	Objets []Objet
}

func (o *Objet) Transsubstancie(typeDeDestination string) *Espece {
	return &Espece{Nom: o.Nom, Type: typeDeDestination}
}

func RitesInitiaux(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	ministres := make([]Humain, 0, 3)
	ministres = append(ministres, celebrant, acolyte1, acolyte2)
	MarcheDesMinistres(ministres, "l'autel")

	celebrant.Dit("Prions")
}

func LiturgieDeLaParole(celebrant *Pretre, lecteur Humain) {
	lecteur.Lit("Première Épitre aux Corinthiens, 1.5")
}

func LiturgieDeLeucharistie(celebrant *Pretre, acolyte1 *Acolyte, acolyte2 *Acolyte) {
	pale := &Objet{Nom: "la pale"}
	corporal := &Objet{Nom: "le corporal"}
	purificatoire := &Objet{Nom: "le purificatoire"}

	lingesLiturgiques := &Objets{Nom: "les linges liturgiques", Objets: []Objet{ *pale, *corporal, *purificatoire }}

	patene := &Objet{Nom: "la patène"}
	calice := &Objet{Nom: "le calice"}

	vasesSacres := &Objets{Nom: "les vases sacrés", Objets: []Objet{ *patene, *calice }}

	acolyte1.MetSurAutel(*vasesSacres)
	acolyte1.MetSurAutel(*lingesLiturgiques)

	pain := &Objet{Nom: "le pain"}
	vin := &Objet{Nom: "le vin"}

	MarcheDesMinistres([]Humain{acolyte1, acolyte2, celebrant}, "l'autel")

	acolyte1.Donne(&Objet{Nom: "le calice"}, celebrant)
	acolyte2.Donne(&Objet{Nom: "le ciboire"}, celebrant)

	acolyte2.Donne(pain, celebrant)
	acolyte2.Donne(vin, celebrant)

	celebrant.Consacre(vin, Sang)
	celebrant.Consacre(pain, Corps)
}

func RitesDeConclusion(celebrant *Pretre, peuple *Peuple) {
	celebrant.Dit("Le Seigneur soit avec vous")
	peuple.Dit("Et avec votre esprit")
	celebrant.Dit("Que Dieu tout-puissant vous bénisse")
	celebrant.Dit("Le Père, le Fils et le Saint-Esprit")
	peuple.Dit("Amen")
	celebrant.Dit("Allez, dans la paix du Christ")
	peuple.Dit("Nous rendons grâce à Dieu")
}

func main() {
	celebrant := &Pretre{Nom: "Claude"}
	acolyte1 := &Acolyte{Nom: "Julien"}
	acolyte2 := &Acolyte{Nom: "Pierre"}
	peuple := &Peuple{Nom: "Le peuple"}

	RitesInitiaux(celebrant, acolyte1, acolyte2)

	LiturgieDeLaParole(celebrant, acolyte1)
	LiturgieDeLeucharistie(celebrant, acolyte1, acolyte2)

	RitesDeConclusion(celebrant, peuple)
}
