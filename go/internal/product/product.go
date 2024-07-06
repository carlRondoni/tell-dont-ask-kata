package product

type Product struct {
	name     string
	category ProductCategory
	price    float64
}

func NewProduct() Product {
	return Product{}
}

func (p Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string) *Product {
	p.name = name
	return p
}

func (p Product) GetCategory() ProductCategory {
	return p.category
}

func (p *Product) SetCategory(category ProductCategory) *Product {
	p.category = category
	return p
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) *Product {
	p.price = price
	return p
}
