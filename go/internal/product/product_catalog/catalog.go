package catalog

import (
	"github.com/carlRondoni/tell-dont-ask-kata/internal/product"
)

type ProductCatalog interface {
	GetByName(name string) (*product.Product, error)
}
