package test_pkg

import (
	"errors"
	"sync"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/product"
)

type InMemoryProductCatalog struct {
	products []*product.Product
	mu       sync.Mutex
}

func NewInMemoryProductCatalog() *InMemoryProductCatalog {
	var products []*product.Product
	food := product.ProductCategory{}
	food.SetName("food")
	food.SetId(10)

	product1 := product.NewProduct()
	product1.SetName("salad")
	product1.SetPrice(3.56)
	product1.SetCategory(food)
	products = append(products, &product1)

	product2 := product.NewProduct()
	product2.SetName("tomato")
	product2.SetPrice(4.56)
	product2.SetCategory(food)
	products = append(products, &product2)

	return &InMemoryProductCatalog{
		products: products,
	}
}

func (c *InMemoryProductCatalog) GetByName(name string) (*product.Product, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, p := range c.products {
		if p.GetName() == name {
			return p, nil
		}
	}

	return nil, errors.New("product not found")
}
