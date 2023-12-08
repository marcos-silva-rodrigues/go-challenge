package entity

import (
	"strings"
)

type Pessoa struct {
	Nome      string
	Idade     int
	Pontuacao int
}

type PorNomeIdade []Pessoa

func (a PorNomeIdade) Len() int { return len(a) }

func (a PorNomeIdade) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a PorNomeIdade) Less(i, j int) bool {
	primeiro := a[i].Nome
	segundo := a[j].Nome

	if strings.ToUpper(primeiro) == strings.ToUpper(segundo) {
		if primeiro == segundo {
			return a[i].Idade < a[j].Idade
		} else {
			return primeiro < segundo
		}
	} else {
		return strings.ToUpper(primeiro) < strings.ToUpper(segundo)
	}

}

func CriaPessoa(nome string, idade, pontuacao int) Pessoa {
	return Pessoa{
		Nome:      nome,
		Idade:     idade,
		Pontuacao: pontuacao,
	}
}
