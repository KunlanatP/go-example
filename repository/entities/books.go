package entities

import (
	"github.com/kunlanat/go-example/domain"
	"github.com/kunlanat/go-example/repository/entities/base"
)

type Books struct {
	base.Model
	Name  string  `json:"name"`
	Desc  string  `json:"description"`
	Price float32 `json:"price"`
}

func (b *Books) ToDomain() *domain.Books {
	return &domain.Books{
		ID:    b.ID,
		Name:  b.Name,
		Desc:  b.Desc,
		Price: b.Price,
	}
}
