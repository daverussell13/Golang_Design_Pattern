package builder

type HouseDirector struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *HouseDirector {
	return &HouseDirector{
		builder,
	}
}

func (d *HouseDirector) SetBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *HouseDirector) BuildHouse() (*House, error) {
	house, err := d.builder.getHouse()
	return house, err
}
