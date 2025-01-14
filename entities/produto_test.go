package entities

import (
	"github.com/govalues/decimal"
	"reflect"
	"testing"
)

func TestProdutoNew(t *testing.T) {
	type args struct {
		identificacao string
		nome          string
		preco         string
	}
	tests := []struct {
		name    string
		args    args
		want    *Produto
		wantErr bool
	}{
		{
			"Funcionando",
			args{
				identificacao: "1234",
				nome:          "banana",
				preco:         "10",
			},
			&Produto{
				Identificacao: "1234",
				Nome:          "banana",
				Preco:         decimal.MustParse("10"),
			},
			false,
		},
		{
			"Nome Invalido",
			args{
				identificacao: "",
				nome:          "banana",
				preco:         "10",
			},
			nil,
			true,
		},
		{
			"Funcionando",
			args{
				identificacao: "1234",
				nome:          "banana",
				preco:         "",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProdutoNew(tt.args.identificacao, tt.args.nome, tt.args.preco)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProdutoNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProdutoNew() got = %v, want %v", got, tt.want)
			}
		})
	}
}
