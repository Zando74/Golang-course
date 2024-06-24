package main

/*
	dependency inversion principle
	-> Les modules de haut niveau ne devraient pas dépendre des modules de bas niveau.
	-> Les deux devraient dépendre d'abstractions.
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// ...
}

type Info struct {
	from         *Person
	Relationship Relationship
	to           *Person
}

// low-level module
type Relationships struct {
	relations []Info
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP
	relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.Relationship == Parent {
			println("John has a child called ", rel.to.name)
		}
	}
}

type ResearchV2 struct {
	browser RelationshipBrowser
}

func (r *ResearchV2) Investigate(name string) {
	for _, p := range r.browser.FindAllChildrenOf(name) {
		println("John has a child called ", p.name)
	}
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for _, rel := range rs.relations {
		if rel.from.name == name && rel.Relationship == Parent {
			result = append(result, rel.to)
		}
	}

	return result
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{relationships}
	research.Investigate()

	researchV2 := ResearchV2{&relationships}
	researchV2.Investigate("John")
}
