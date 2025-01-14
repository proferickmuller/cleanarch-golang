package usecases

import (
	"ca-go/entities"
	"ca-go/interfaces"
	"errors"
)

type ProdutoObterPorIdUseCase struct {
	produtoGateway interfaces.IProdutoGateway
}

func (u *ProdutoObterPorIdUseCase) Run(identificacao string) (*entities.Produto, error) {
	produto, err := u.produtoGateway.BuscarPorIdentificacao(identificacao)

	if err != nil {
		return nil, errors.New("erro ao buscar: " + err.Error())
	}

	return produto, nil
}
