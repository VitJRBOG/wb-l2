package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

func ExecuteBuilderExample() {
	aDirector := Director{}

	masterBuilder := GetBuilder("cottage")
	if masterBuilder == nil {
		panic(fmt.Errorf("can't build the 'cottage'"))
	}

	aDirector.SetBuilder(masterBuilder)
	house := aDirector.BuildTheHouse()
	fmt.Println(house)

	masterBuilder = GetBuilder("townhouse")
	if masterBuilder == nil {
		panic(fmt.Errorf("can't build the 'townhouse'"))
	}

	aDirector.SetBuilder(masterBuilder)
	house = aDirector.BuildTheHouse()
	fmt.Println(house)

	masterBuilder = GetBuilder("hut")
	if masterBuilder == nil {
		panic(fmt.Errorf("can't build the 'hut'"))
	}
}

type Director struct {
	MasterBuilder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.MasterBuilder = b
}

func (d *Director) BuildTheHouse() *House {
	d.MasterBuilder.SetFloorType()
	d.MasterBuilder.SetFloorNumbers()
	d.MasterBuilder.SetWallsType()
	d.MasterBuilder.SetWindowsType()
	d.MasterBuilder.SetRoofType()

	return d.MasterBuilder.GetHouse()
}

type Builder interface {
	GetHouse() *House
	SetFloorType()
	SetFloorNumbers()
	SetWallsType()
	SetWindowsType()
	SetRoofType()
}

func GetBuilder(buildingType string) Builder {
	switch buildingType {
	case "cottage":
		return NewCottage()
	case "townhouse":
		return NewTownHouse()
	default:
		return nil
	}
}

type House struct {
	FloorType    string
	FloorNumbers int
	WallsType    string
	WindowsType  string
	RoofType     string
}

type Cottage struct {
	FloorType    string
	FloorNumbers int
	WallsType    string
	WindowsType  string
	RoofType     string
}

func NewCottage() *Cottage {
	return &Cottage{}
}

func (c *Cottage) GetHouse() *House {
	return &House{
		FloorType:    c.FloorType,
		FloorNumbers: c.FloorNumbers,
		WallsType:    c.WallsType,
		WindowsType:  c.WindowsType,
		RoofType:     c.RoofType,
	}
}

func (c *Cottage) SetFloorType() {
	c.FloorType = "wood"
	fmt.Printf("Floor type will be: %s\n", c.FloorType)
}

func (c *Cottage) SetFloorNumbers() {
	c.FloorNumbers = 1
	fmt.Printf("Floor numbers will be: %d\n", c.FloorNumbers)
}

func (c *Cottage) SetWallsType() {
	c.WallsType = "brick"
	fmt.Printf("Walls type will be: %s\n", c.WallsType)
}

func (c *Cottage) SetWindowsType() {
	c.WindowsType = "oval"
	fmt.Printf("Windows type will be: %s\n", c.WindowsType)
}

func (c *Cottage) SetRoofType() {
	c.RoofType = "gambrel"
	fmt.Printf("Roof type will be: %s\n", c.RoofType)
}

type TownHouse struct {
	FloorType    string
	FloorNumbers int
	WallsType    string
	WindowsType  string
	RoofType     string
}

func NewTownHouse() *TownHouse {
	return &TownHouse{}
}

func (t *TownHouse) GetHouse() *House {
	return &House{
		FloorType:    t.FloorType,
		FloorNumbers: t.FloorNumbers,
		WallsType:    t.WallsType,
		WindowsType:  t.WindowsType,
		RoofType:     t.RoofType,
	}
}

func (t *TownHouse) SetFloorType() {
	t.FloorType = "concrete"
	fmt.Printf("Floor type will be: %s\n", t.FloorType)
}

func (t *TownHouse) SetFloorNumbers() {
	t.FloorNumbers = 3
	fmt.Printf("Floor numbers will be: %d\n", t.FloorNumbers)
}

func (t *TownHouse) SetWallsType() {
	t.WallsType = "drywall"
	fmt.Printf("Walls type will be: %s\n", t.WallsType)
}

func (t *TownHouse) SetWindowsType() {
	t.WindowsType = "rectangle"
	fmt.Printf("Windows type will be: %s\n", t.WindowsType)
}

func (t *TownHouse) SetRoofType() {
	t.RoofType = "gable"
	fmt.Printf("Roof type will be: %s\n", t.RoofType)
}
