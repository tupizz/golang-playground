package main

import "fmt"

type Pessoa struct{}

type PessoaWithStr struct {
	Name string
	Job  string
}

// To String way to do
//
// method from the struct respects the interface fmt.Stringer
//
//	type Stringer interface {
//		String() string
//	}
func (p PessoaWithStr) String() string {
	return fmt.Sprintf("Name here: %s, Job here: %s", p.Name, p.Job)
}

func main() {
	p := Pessoa{}
	fmt.Println(p)

	pws := PessoaWithStr{
		Name: "Tadeu",
		Job:  "Software Engineer",
	}
	fmt.Println(pws)

}
