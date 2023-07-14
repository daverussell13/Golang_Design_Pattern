package informa

type BeanBag struct{}

func (b BeanBag) Price() float64 {
	return 1199000
}

func (b BeanBag) IsIoTEnabled() bool {
	return false
}

func (b BeanBag) IsSoft() bool {
	return true
}
