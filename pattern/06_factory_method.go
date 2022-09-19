package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

func ExecuteFactoryMethodExample() {
	f := Fabric{}

	sizes := []string{"small", "medium", "big"}

	for _, gunSize := range sizes {
		g := f.GetGun(gunSize)
		if g == nil {
			panic("unknown gun size")
		}
		g.Shoot()
	}
}

type Fabric struct{}

func (f *Fabric) GetGun(gunSize string) Gun {
	switch gunSize {
	case "small":
		return newPistol()
	case "medium":
		return newMachinegun()
	case "big":
		return newRocketlauncher()
	default:
		return nil
	}
}

type Gun interface {
	Shoot()
}

type pistol struct {
	sound string
}

func newPistol() *pistol {
	return &pistol{"piy-piy"}
}

func (p *pistol) Shoot() {
	fmt.Printf("Pistol does %s\n", p.sound)
}

type machinegun struct {
	sound string
}

func newMachinegun() *machinegun {
	return &machinegun{"ratata"}
}

func (m *machinegun) Shoot() {
	fmt.Printf("Machinegun does %s\n", m.sound)
}

type rocketlauncher struct {
	sound string
}

func newRocketlauncher() *rocketlauncher {
	return &rocketlauncher{"kaboom"}
}

func (r *rocketlauncher) Shoot() {
	fmt.Printf("Rocketlauncher does %s\n", r.sound)
}
