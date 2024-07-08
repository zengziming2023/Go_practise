package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name
	Id      int
	Name    string
	Origin  []string
}

func (p *Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v\n", p.Id, p.Name, p.Origin)
}

type Nesting struct {
	XMLName xml.Name
	Plants  []*Plant
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"} //, XMLName: xml.Name{
	//Space: "XML SPACE",
	//	Local: "CN",
	//}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, err := xml.MarshalIndent(coffee, "  ", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p Plant
	err = xml.Unmarshal(out, &p)
	if err != nil {
		return
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, err = xml.MarshalIndent(nesting, " ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("nesting:", string(out))
	fmt.Println(xml.Header + string(out))
}
