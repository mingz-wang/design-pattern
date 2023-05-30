package mediator

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"

	fmt.Printf("CD Driver: reading data %s\n", c.Data)

	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: spit data with Sound %s, Video %s\n", c.Sound, c.Video)

	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (vc *VideoCard) Display(data string) {
	vc.Data = data
	fmt.Printf("VideoCard: display %s\n", vc.Data)
	GetMediatorInstance().changed(vc)
}

type SoundCard struct {
	Data string
}

func (sc *SoundCard) Play(data string) {
	sc.Data = data
	fmt.Printf("SoundCard: display %s\n", sc.Data)
	GetMediatorInstance().changed(sc)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

func (m *Mediator) changed(i any) {
	switch instance := i.(type) {
	case *CDDriver:
		m.CPU.Process(instance.Data)
	case *CPU:
		m.Sound.Play(instance.Sound)
		m.Video.Display(instance.Video)
	}
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}
