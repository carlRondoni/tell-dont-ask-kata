package product

type ProductCategory struct {
	name          string
	taxPercentage float64
}

func (c ProductCategory) GetName() string {
	return c.name
}

func (c ProductCategory) GetTaxPercentage() float64 {
	return c.taxPercentage
}

func (c *ProductCategory) SetName(name string) {
	c.name = name
}

func (c *ProductCategory) SetId(taxPercentage float64) {
	c.taxPercentage = taxPercentage
}
