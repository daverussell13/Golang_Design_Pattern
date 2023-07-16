package builder

type HouseBuilder interface {
	SetHouseAddress(string) HouseBuilder
	SetWindowType(string) HouseBuilder
	SetDoorType(string) HouseBuilder
	SetNumFloor(int) HouseBuilder
	getHouse() (*House, error)
}

func GetBuilder(builderType string) HouseBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
