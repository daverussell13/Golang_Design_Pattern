package factory

import "fmt"

type Content interface {
	Play()
}

type TwitchContent struct{}

func (c *TwitchContent) Play() {
	fmt.Println("Playing twitch content")
}

type YoutubeContent struct{}

func (p *YoutubeContent) Play() {
	fmt.Println("Playing youtube content")
}
