package foo

type Profissao struct {
	Name  string `json:"profissao_name"`
	Place string `json:"place"`
}

func (p *Profissao) UpdateProfissaoName(name string) {
	p.Name = name
}
