package usecases

import (
	"ca-go/entities"
	"ca-go/interfaces"
	"errors"
)

type ProdutoIncluirUseCase struct {
	produtoGateway interfaces.IProdutoGateway
}

func (p *ProdutoIncluirUseCase) Run(identificacao string, nome string, preco string) (*entities.Produto, error) {
	produto, err := entities.ProdutoNew(identificacao, nome, preco)
	if err != nil {
		return nil, errors.New("Criação de produto inválida: " + err.Error())
	}

	err = p.produtoGateway.Adicionar(produto)
	if err != nil {
		return nil, errors.New("Criação de produto inválida: " + err.Error())
	}

	return produto, nil
}
