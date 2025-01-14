package interfaces

import "ca-go/entities"

type IProdutoGateway interface {
	Adicionar(*entities.Produto) error
	BuscarPorNome(nome string, exato bool) ([]*entities.Produto, error)
	BuscarPorIdentificacao(identificacao string) (*entities.Produto, error)
}
