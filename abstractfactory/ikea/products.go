package ikea

import "Design_Pattern/abstractfactory"

type Leifarne struct{}

func (l Leifarne) Price() float64 {
	return 10950000
}

func (l Leifarne) IsIoTEnabled() bool {
	return false
}

func (l Leifarne) IsSoft() bool {
	return false
}

type Vittsjo struct{}

func (v Vittsjo) Price() float64 {
	return 899000
}

func (v Vittsjo) Size() abstractfactory.Dimension {
	return abstractfactory.Dimension{
		Length: 80,
		Width:  78,
		Height: 40,
	}
}

func (v Vittsjo) IsFoldable() bool {
	return false
}

type Hemlingby struct{}

func (h Hemlingby) Price() float64 {
	return 1795000
}

func (h Hemlingby) Style() abstractfactory.SofaStyle {
	return abstractfactory.AmericanStyle
}
