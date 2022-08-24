package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

func ExecuteFactoryMethodExample() {
	ar := GetGun("assault_rifle")
	if ar != nil {
		fmt.Printf("This assault rifle is an '%s'\n", ar.GetName())
	} else {
		panic(fmt.Errorf("unknown gun: %s", "assault_rifle"))
	}

	sr := GetGun("sniper_rifle")
	if sr != nil {
		fmt.Printf("This sniper rifle is an '%s'\n", sr.GetName())
	} else {
		panic(fmt.Errorf("unknown gun: %s", "sniper_rifle"))
	}

	mg := GetGun("machine_gun")
	if mg != nil {
		fmt.Printf("This machine gun is an '%s'\n", mg.GetName())
	} else {
		panic(fmt.Errorf("unknown gun: %s", "machine_gun"))
	}

	rl := GetGun("rocket_launcher")
	if rl != nil {
		fmt.Printf("This rocket launcher is an '%s'\n", rl.GetName())
	} else {
		panic(fmt.Errorf("unknown gun: %s", "rocket_launcher"))
	}
}

func GetGun(gunType string) Armorer {
	switch gunType {
	case "assault_rifle":
		return NewAssaultRifle()
	case "sniper_rifle":
		return NewSniperRifle()
	case "machine_gun":
		return NewMachineGun()
	default:
		return nil
	}
}

type Armorer interface {
	SetName(name string)
	GetName() string
}

type Gun struct {
	name string
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

type AssaultRifle struct {
	Gun
}

func NewAssaultRifle() *AssaultRifle {
	return &AssaultRifle{
		Gun: Gun{
			name: "M4",
		},
	}
}

//

type SniperRifle struct {
	Gun
}

func NewSniperRifle() *SniperRifle {
	return &SniperRifle{
		Gun: Gun{
			name: "PGM Hecate II",
		},
	}
}

//

type MachineGun struct {
	Gun
}

func NewMachineGun() *MachineGun {
	return &MachineGun{
		Gun: Gun{
			name: "RPK",
		},
	}
}
