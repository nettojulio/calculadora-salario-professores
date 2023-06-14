package models

import (
	"fmt"
)

type ProfessorHorista struct {
	Professor
	ValorHora        float64
	HorasTrabalhadas int64
}

func (h *ProfessorHorista) CalcularSalario() float64 {
	return h.ValorHora * float64(h.HorasTrabalhadas)
}

func (h *ProfessorHorista) Display() string {
	return fmt.Sprintf("O salário do professor %s %s é: R$ %.2f", h.Professor.GetTipoContrato(), h.Professor.GetNome(), h.CalcularSalario())
}

func NewProfessorHorista(nome string, valorHora float64, horasTrabalhadas int64) *ProfessorHorista {
	return &ProfessorHorista{
		Professor: Professor{
			Nome:         nome,
			TipoContrato: 2,
		},
		ValorHora:        valorHora,
		HorasTrabalhadas: horasTrabalhadas,
	}
}
