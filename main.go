package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/marcos-silva-rodrigues/go-challenge/entity"
)

func main() {
	origem := os.Args[1]
	destino := os.Args[2]

	arquivo := abreArquivo(origem)
	registros := converteCsv(arquivo)

	cabecalhos := registros[0]

	entidadeEmlinha := registros[1:]
	pessoas := converteLinhasParaEntidades(entidadeEmlinha)

	sort.Sort(entity.PorNomeIdade(pessoas))

	w := criaArquivoCsvDeDestino(destino)
	w.Write(cabecalhos)

	preencheCsv(w, pessoas)
}

func abreArquivo(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func converteCsv(file *os.File) [][]string {
	r := csv.NewReader(file)
	registros, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
	return registros
}

func converteLinhasParaEntidades(registros [][]string) []entity.Pessoa {
	var listaPessoa []entity.Pessoa

	// Pular os cabeçalhos
	for _, linha := range registros {
		nome := linha[0]
		idade, _ := strconv.Atoi(linha[1])
		pontuacao, _ := strconv.Atoi(linha[2])

		pessoa := entity.CriaPessoa(nome, idade, pontuacao)
		listaPessoa = append(listaPessoa, pessoa)
	}

	return listaPessoa
}

func criaArquivoCsvDeDestino(filename string) *csv.Writer {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	return writer
}

func preencheCsv(w *csv.Writer, pessoas []entity.Pessoa) {
	for _, pessoa := range pessoas {
		data := []string{
			pessoa.Nome, strconv.Itoa(pessoa.Idade), strconv.Itoa(pessoa.Pontuacao),
		}
		if err := w.Write(data); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
