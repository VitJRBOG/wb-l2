package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

func ExecuteCommandExample() {
	onCmd := &OnCommand{}
	offCmd := &OffCommand{}

	onBtn := &OnButton{onCmd}
	offBtn := &OffButton{offCmd}

	computer := Computer{onBtn: onBtn, offBtn: offBtn}

	computer.On()
	computer.Off()
}

type Computer struct {
	onBtn  Button
	offBtn Button
}

func (c *Computer) On() {
	c.onBtn.Press()
}

func (c *Computer) Off() {
	c.offBtn.Press()
}

type Button interface {
	Press()
}

type OnButton struct {
	onCommand Command
}

func (b *OnButton) Press() {
	b.onCommand.Exec()
}

type OffButton struct {
	offCommand Command
}

func (b *OffButton) Press() {
	b.offCommand.Exec()
}

type Command interface {
	Exec()
}

type OnCommand struct{}

func (c *OnCommand) Exec() {
	fmt.Println("Turning on")
}

type OffCommand struct{}

func (c *OffCommand) Exec() {
	fmt.Println("Turning off")
}
