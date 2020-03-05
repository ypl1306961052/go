package main

type Person struct {
	name string
	age  int
	city string
}
type PersonService struct {
	Persons []Person
}

func (ps *PersonService) findPersonByName(name string, persons *[]Person) error {

	var persTmp = make([]Person, 10)
	for _, p := range ps.Persons {
		if p.name == name {
			persTmp = append(persTmp, p)
		}

	}
	persons = &persTmp
	return nil

}
//


func main() {
	rpcPersonService:=new(PersonService)
}
