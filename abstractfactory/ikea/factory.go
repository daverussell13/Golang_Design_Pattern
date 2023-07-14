package ikea

import "Design_Pattern/abstractfactory"

type Ikea struct{}

func (i Ikea) CreateChair() abstractfactory.Chair {
	return &Leifarne{}
}

func (i Ikea) CreateTable() abstractfactory.Table {
	return &Vittsjo{}
}

func (i Ikea) CreateSofa() abstractfactory.Sofa {
	return &Hemlingby{}
}
