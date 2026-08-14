package motherboard

import "fmt"

// aqui injetaria a dependencia de banco de dados
type MotherboardRepository struct{}

type Motherboard struct {
	Name  string
	Color string
	Price float64
}

func (r *MotherboardRepository) FindByName(name string) (*Motherboard, error) {
	return &Motherboard{
		Name:  name,
		Color: "black",
		Price: 899.90,
	}, nil
}

func (r *MotherboardRepository) Insert(description string, colorID, manufacturerID int32) (string, error) {
	generatedID := fmt.Sprintf("mb-%d-%d", colorID, manufacturerID)
	return generatedID, nil
}
