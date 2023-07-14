package factory

type ContentCreatorType int

type ContentCreator interface {
	Produce() Content
}

type TwitchContentCreator struct{}

func (t *TwitchContentCreator) Produce() Content {
	return &TwitchContent{}
}

type YoutubeContentCreator struct{}

func (y *YoutubeContentCreator) Produce() Content {
	return &YoutubeContent{}
}
