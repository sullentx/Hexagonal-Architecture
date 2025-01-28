package entities

//SRP, Single Responsability Principle (Principio de responsabilidad única)
type Product struct {
	ID    int
	Name  string
	Price float64
}

func NewProduct(Name string, Price float64) *Product {
	return &Product{Name: Name, Price: Price}
}
