package main

import "fmt"

type Address2 struct {
	StreetAddress2, City, Country string
}

func (a *Address2) DeepCopy() *Address2 {
	return &Address2{
		a.StreetAddress2,
		a.City,
		a.Country}
}

type Person2 struct {
	Name     string
	Address2 *Address2
	Friends  []string
}

func (p *Person2) DeepCopy() *Person2 {
	q := *p // copies Name
	q.Address2 = p.Address2.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main2() {
	john := Person2{"John",
		&Address2{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address2.StreetAddress2 = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address2)
	fmt.Println(jane, jane.Address2)
}
