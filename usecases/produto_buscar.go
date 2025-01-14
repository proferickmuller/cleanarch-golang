package usecases

import (
	"ca-go/entities"
	"ca-go/interfaces"
	"errors"
)

type ProdutoBuscarPorNomeUseCase struct {
	produtoGateway interfaces.IProdutoGateway
}

func (p *ProdutoBuscarPorNomeUseCase) Run(nome string, exato bool) ([]*entities.Produto, error) {
	produtos, err := p.produtoGateway.BuscarPorNome(nome, exato)
	if err != nil {
		return nil, errors.New("erro ao buscar: " + err.Error())
	}

	return produtos, nil
}
