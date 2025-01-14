package entities

import (
	"errors"
	"github.com/govalues/decimal"
	"strings"
)

type Produto struct {
	Identificacao string
	Nome          string
	Preco         decimal.Decimal
}

func ProdutoNew(identificacao string, nome string, preco string) (*Produto, error) {
	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || strings.TrimSpace(preco) == "" {
		return nil, errors.New("Nenhum dos campos pode estar em branco")
	}

	d, err := decimal.Parse(preco)

	if err != nil {
		return nil, errors.New("Preco deve ser um número válido")
	}

	return &Produto{
		Identificacao: identificacao,
		Nome:          nome,
		Preco:         d,
	}, nil

}
