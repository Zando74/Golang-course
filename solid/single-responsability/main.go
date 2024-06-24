package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	Une seule responsabilité
	-> Une seule raison de changer
	-> Une seule raison d'exister
	-> Separation des préoccupations

	Attention anti-pattern : God Object
*/

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// attention, ces fonctions de devraient pas être ici.

func (j *Journal) Save(filename string) {
}

func (j *Journal) Load(filename string) {
}

func (j *Journal) LoadFromWeb(url string) {
}

// La responsabilité de sauvegarder est donnée à un autre objet
type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	fmt.Println(j.String())

	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
