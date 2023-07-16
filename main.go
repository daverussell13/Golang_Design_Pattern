package main

import (
	"Design_Pattern/abstractfactory"
	"Design_Pattern/abstractfactory/ikea"
	"Design_Pattern/abstractfactory/informa"
	"Design_Pattern/builder"
	"Design_Pattern/factory"
	"fmt"
	"log"
)

// Factory method = ContentCreator
func factoryDp() {
	var contentCreator factory.ContentCreator
	var content factory.Content

	contentCreator = &factory.TwitchContentCreator{}
	content = contentCreator.Produce()
	content.Play()

	contentCreator = &factory.YoutubeContentCreator{}
	content = contentCreator.Produce()
	content.Play()
}

// Abstract factory method = Furniture factory, Factory method = Chair, Table, Sofa
func abstractFactoryDp() {
	var furnitureFactory abstractfactory.FurnitureFactory

	furnitureFactory = &ikea.Ikea{}
	fmt.Println(furnitureFactory.CreateChair().Price())

	furnitureFactory = &informa.Informa{}
	fmt.Println(furnitureFactory.CreateChair().Price())
}

func builderDp() {
	normalBuilder := builder.GetBuilder("normal").
		SetHouseAddress("Somewhere in your heart") // if you build with NormalHouseBuilder you need to set the address

	iglooBuilder := builder.GetBuilder("igloo").
		SetDoorType("Snow Iron Door").
		SetWindowType("Snow Iron Window")

	director := builder.NewDirector(normalBuilder)
	if normalHouse, err := director.BuildHouse(); err != nil {
		log.Println("Error occurred when building house")
		log.Println(err)
	} else {
		fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
		fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
		fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)
	}

	// Changing builder
	director.SetBuilder(iglooBuilder)
	if iglooHouse, err := director.BuildHouse(); err != nil {
		log.Println("Error occurred when building house")
		log.Println(err)
	} else {
		fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
		fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
		fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
	}
}

func main() {
	/* Uncomment if you want to test */
	//factoryDp()
	//abstractFactoryDp()
	builderDp()
}
