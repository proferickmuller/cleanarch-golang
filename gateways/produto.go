package gateways

import (
	"ca-go/dto"
	"ca-go/entities"
	"ca-go/interfaces"
	"github.com/govalues/decimal"
)

type ProdutoGateway struct {
	dataSource interfaces.IDataSource
}

func (p *ProdutoGateway) Adicionar(produto *entities.Produto) error {
	prod_dto := &dto.ProdutoDTO{
		Identificacao: produto.Identificacao,
		Nome:          produto.Nome,
		Preco:         produto.Preco.String(),
	}
	err := p.dataSource.ProdutoInsert(prod_dto)
	return err
}

func (p *ProdutoGateway) BuscarPorNome(nome string, exato bool) ([]*entities.Produto, error) {
	prods_dto, err := p.dataSource.ProdutoByName(nome, exato)
	if err != nil {
		return nil, err
	}

	var ents []*entities.Produto
	for _, item := range prods_dto {
		ents = append(ents, &entities.Produto{
			Identificacao: item.Identificacao,
			Nome:          item.Nome,
			Preco:         decimal.MustParse(item.Preco),
		})
	}

	return ents, nil

}

func (p *ProdutoGateway) BuscarPorIdentificacao(identificacao string) (*entities.Produto, error) {
	prod_dto, err := p.dataSource.ProdutoById(identificacao)
	if err != nil {
		return nil, err
	}

	if prod_dto != nil {
		return &entities.Produto{
			Identificacao: prod_dto.Identificacao,
			Nome:          prod_dto.Nome,
			Preco:         decimal.MustParse(prod_dto.Preco),
		}, nil
	} else {
		return nil, nil
	}
}
