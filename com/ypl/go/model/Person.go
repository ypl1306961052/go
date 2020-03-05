package model

import "fmt"

type PrintAdder interface {
	PrintAdder()
}

func init() {
	fmt.Println("初始化")
}

type Adder struct {
	City string `json:"city"`
	//
	Province string `json:"province"`
}
type Man struct {
	Name string `json:"name"`
}
type Person struct {
	Man
	Age   uint8 `json:"age"`
	Adder Adder `json:"adder"`
}

func (m *Man) SayName() {
	fmt.Println("名字为:", m.Name)
}
func (p *Person) PrintAdder() {
	fmt.Println(p.Adder.City, p.Adder.Province)
}
