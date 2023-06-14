package models

import "calculadora-salario-professores/internal/tools"

type ImplProfessor interface {
	GetNome() string
	GetTipoContrato() tools.Modelos
	CalcularSalario() float64
	Display() string
}

type Professor struct {
	Nome         string
	TipoContrato tools.Modelos
}

type ProfessorDTO struct {
	Id, IdDetalhes     int64
	Nome, TipoContrato string
	Inativo            bool
}

func (p *Professor) GetNome() string {
	return p.Nome
}

func (p *Professor) GetTipoContrato() tools.Modelos {
	return p.TipoContrato
}

func NewProfessorDTO(id int64, nome string, tipoContrato string, idDetalhes int64, inativo bool) ProfessorDTO {
	return ProfessorDTO{
		Id:           id,
		Nome:         nome,
		TipoContrato: tipoContrato,
		IdDetalhes:   idDetalhes,
		Inativo:      inativo,
	}
}
