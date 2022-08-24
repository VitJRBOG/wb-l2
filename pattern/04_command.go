package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

func ExecuteCommandExample() {
	computer := Computer{}

	onCommand := OnCommand{
		device: &computer,
	}

	offCommand := OffCommand{
		device: &computer,
	}

	onButton := Button{
		command: &onCommand,
	}

	offButton := Button{
		command: &offCommand,
	}

	onButton.Press()
	offButton.Press()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

type Command interface {
	Execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

type Device interface {
	On()
	Off()
}

type Computer struct {
	status string
}

func (c *Computer) On() {
	c.status = "on"
	fmt.Println("Computer is turned on now")
}

func (c *Computer) Off() {
	c.status = "off"
	fmt.Println("Computer is turned off now")
}
