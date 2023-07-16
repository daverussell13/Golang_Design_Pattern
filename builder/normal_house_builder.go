package builder

import "errors"

type NormalHouseBuilder struct {
	address    string
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalHouseBuilder {
	return &NormalHouseBuilder{
		windowType: "Wooden Window",
		doorType:   "Wooden Door",
		floor:      2,
	}
}

func (b *NormalHouseBuilder) SetWindowType(windowType string) HouseBuilder {
	b.windowType = windowType
	return b
}

func (b *NormalHouseBuilder) SetDoorType(doorType string) HouseBuilder {
	b.doorType = doorType
	return b
}

func (b *NormalHouseBuilder) SetNumFloor(numFloor int) HouseBuilder {
	b.floor = numFloor
	return b
}

func (b *NormalHouseBuilder) SetHouseAddress(address string) HouseBuilder {
	b.address = address
	return b
}

func (b *NormalHouseBuilder) getHouse() (*House, error) {
	if b.address == "" {
		return nil, errors.New("house need to have address")
	}
	return &House{
		Address:    b.address,
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}, nil
}
