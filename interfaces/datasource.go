package interfaces

import "ca-go/dto"

type IDataSource interface {
	ProdutoInsert(dto *dto.ProdutoDTO) error
	ProdutoById(id string) (*dto.ProdutoDTO, error)
	ProdutoByName(name string, exact bool) ([]*dto.ProdutoDTO, error)
}
