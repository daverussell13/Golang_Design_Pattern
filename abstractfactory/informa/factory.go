package informa

import "Design_Pattern/abstractfactory"

type Informa struct{}

func (i Informa) CreateChair() abstractfactory.Chair {
	return &BeanBag{}
}

func (i Informa) CreateTable() abstractfactory.Table {
	return nil
}

func (i Informa) CreateSofa() abstractfactory.Sofa {
	return nil
}
