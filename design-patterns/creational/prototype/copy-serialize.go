package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address3 struct {
	StreetAddress3, City, Country string
}

type Person3 struct {
	Name     string
	Address3 *Address3
	Friends  []string
}

func (p *Person3) DeepCopy() *Person3 {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person3{}
	_ = d.Decode(&result)
	return &result
}

func main3() {
	john := Person3{"John",
		&Address3{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt", "Sam"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address3.StreetAddress3 = "321 Baker St"
	jane.Friends = append(jane.Friends, "Jill")

	fmt.Println(john, john.Address3)
	fmt.Println(jane, jane.Address3)

}
