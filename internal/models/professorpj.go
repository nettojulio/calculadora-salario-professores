package models

import (
	"fmt"
)

type ProfessorPJ struct {
	Professor
	Contrato float64
}

func (p *ProfessorPJ) CalcularSalario() float64 {
	return p.Contrato
}

func (p *ProfessorPJ) Display() string {
	return fmt.Sprintf("O salário do professor %s %s é: R$ %.2f", p.Professor.GetTipoContrato(), p.Professor.GetNome(), p.CalcularSalario())
}

func NewProfessorPJ(nome string, contrato float64) *ProfessorPJ {
	return &ProfessorPJ{
		Professor: Professor{
			Nome:         nome,
			TipoContrato: 3,
		},
		Contrato: contrato,
	}
}
