package builder

type IglooHouseBuilder struct {
	address    string
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooHouseBuilder {
	return &IglooHouseBuilder{
		address:    "North Pole",
		windowType: "Snow Window",
		doorType:   "Snow Door",
	}
}

func (b *IglooHouseBuilder) SetHouseAddress(address string) HouseBuilder {
	b.address = address
	return b
}

func (b *IglooHouseBuilder) SetWindowType(windowType string) HouseBuilder {
	b.windowType = windowType
	return b
}

func (b *IglooHouseBuilder) SetDoorType(doorType string) HouseBuilder {
	b.doorType = doorType
	return b
}

func (b *IglooHouseBuilder) SetNumFloor(numFloor int) HouseBuilder {
	b.floor = numFloor
	return b
}

func (b *IglooHouseBuilder) getHouse() (*House, error) {
	return &House{
		Address:    b.address,
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}, nil
}
