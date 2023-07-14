package abstractfactory

type Pricing interface {
	Price() float64
}

type Chair interface {
	Pricing
	IsIoTEnabled() bool
	IsSoft() bool
}

type Dimension struct {
	Length, Width, Height int
}

type Table interface {
	Pricing
	Size() Dimension
	IsFoldable() bool
}

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "European"
	AmericanStyle SofaStyle = "American"
)

type Sofa interface {
	Pricing
	Style() SofaStyle
}
