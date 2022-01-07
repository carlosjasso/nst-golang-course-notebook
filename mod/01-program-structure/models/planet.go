package models

/*
Glossary:
struct: A struct is a user-defined type that contains a collection of named fields/properties.
lowercase identifiers - are only accesible within the package and is not exported
uppercase identifiers - are exported and can be accesed from the the package gets imported
*/

type Planet struct {
	Name       string
	Population uint64
	alias      string
}

/*
struct field seters can be custom-written for Encapsulation
*/
func (p *Planet) SetAlias(alias string) {
	p.alias = alias
}
