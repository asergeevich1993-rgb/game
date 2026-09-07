package equipment

import "go-from-zero/concurrency3/miner"

type Equipment struct {
	Name   string
	Price  miner.Coal
	Bought bool
}

func NewEquip(name string) *Equipment {

	switch name {
	case "кирка":
		return &Equipment{
			Name:   "Кирка",
			Price:  3000,
			Bought: false,
		}
	case "Вентиляция":
		return &Equipment{
			Name:   "Вентиляция",
			Price:  15000,
			Bought: false,
		}
	case "Вагонетка":
		return &Equipment{
			Name:   "Вагонетка",
			Price:  50000,
			Bought: false,
		}
	}
	return nil
}
