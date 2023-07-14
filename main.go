package main

import "Design_Pattern/factory"

// The intention here is to make a factory of a content
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

func main() {
	factoryDp()
}
