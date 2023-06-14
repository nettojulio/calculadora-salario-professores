package models

import (
	"fmt"
)

type ProfessorCLT struct {
	Professor
	Salario float64
}

func (c *ProfessorCLT) CalcularSalario() float64 {
	return c.Salario
}

func (c *ProfessorCLT) Display() string {
	return fmt.Sprintf("O salário do professor %s %s é: R$ %.2f", c.Professor.GetTipoContrato(), c.Professor.GetNome(), c.CalcularSalario())

}

func NewProfessorCLT(nome string, salario float64) *ProfessorCLT {
	return &ProfessorCLT{
		Professor: Professor{
			Nome:         nome,
			TipoContrato: 1,
		},
		Salario: salario,
	}
}
