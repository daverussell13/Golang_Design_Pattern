package main

import (
	"Design_Pattern/abstractfactory"
	"Design_Pattern/abstractfactory/ikea"
	"Design_Pattern/abstractfactory/informa"
	"Design_Pattern/factory"
	"fmt"
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

func main() {
	factoryDp()
	abstractFactoryDp()
}
