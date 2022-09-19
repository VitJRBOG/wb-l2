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

	cottageBuilder := NewCottage()

	aDirector.SetBuilder(cottageBuilder)
	house := aDirector.BuildTheHouse()
	fmt.Println(house)

	townhouseBuilder := NewTownHouse()

	aDirector.SetBuilder(townhouseBuilder)
	house = aDirector.BuildTheHouse()
	fmt.Println(house)
}

type Director struct {
	MasterBuilder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.MasterBuilder = b
}

func (d *Director) BuildTheHouse() *House {
	d.MasterBuilder.setFloorType()
	d.MasterBuilder.setFloorNumbers()
	d.MasterBuilder.setWallsType()
	d.MasterBuilder.setWindowsType()
	d.MasterBuilder.setRoofType()

	return d.MasterBuilder.getHouse()
}

type Builder interface {
	getHouse() *House
	setFloorType()
	setFloorNumbers()
	setWallsType()
	setWindowsType()
	setRoofType()
}

type House struct {
	FloorType    string
	FloorNumbers int
	WallsType    string
	WindowsType  string
	RoofType     string
}

type Cottage struct {
	floorType    string
	floorNumbers int
	wallsType    string
	windowsType  string
	roofType     string
}

func NewCottage() *Cottage {
	return &Cottage{}
}

func (c *Cottage) getHouse() *House {
	return &House{
		FloorType:    c.floorType,
		FloorNumbers: c.floorNumbers,
		WallsType:    c.wallsType,
		WindowsType:  c.windowsType,
		RoofType:     c.roofType,
	}
}

func (c *Cottage) setFloorType() {
	c.floorType = "wood"
	fmt.Printf("Floor type will be: %s\n", c.floorType)
}

func (c *Cottage) setFloorNumbers() {
	c.floorNumbers = 1
	fmt.Printf("Floor numbers will be: %d\n", c.floorNumbers)
}

func (c *Cottage) setWallsType() {
	c.wallsType = "brick"
	fmt.Printf("Walls type will be: %s\n", c.wallsType)
}

func (c *Cottage) setWindowsType() {
	c.windowsType = "oval"
	fmt.Printf("Windows type will be: %s\n", c.windowsType)
}

func (c *Cottage) setRoofType() {
	c.roofType = "gambrel"
	fmt.Printf("Roof type will be: %s\n", c.roofType)
}

type TownHouse struct {
	floorType    string
	floorNumbers int
	wallsType    string
	windowsType  string
	roofType     string
}

func NewTownHouse() *TownHouse {
	return &TownHouse{}
}

func (t *TownHouse) getHouse() *House {
	return &House{
		FloorType:    t.floorType,
		FloorNumbers: t.floorNumbers,
		WallsType:    t.wallsType,
		WindowsType:  t.windowsType,
		RoofType:     t.roofType,
	}
}

func (t *TownHouse) setFloorType() {
	t.floorType = "concrete"
	fmt.Printf("Floor type will be: %s\n", t.floorType)
}

func (t *TownHouse) setFloorNumbers() {
	t.floorNumbers = 3
	fmt.Printf("Floor numbers will be: %d\n", t.floorNumbers)
}

func (t *TownHouse) setWallsType() {
	t.wallsType = "drywall"
	fmt.Printf("Walls type will be: %s\n", t.wallsType)
}

func (t *TownHouse) setWindowsType() {
	t.windowsType = "rectangle"
	fmt.Printf("Windows type will be: %s\n", t.windowsType)
}

func (t *TownHouse) setRoofType() {
	t.roofType = "gable"
	fmt.Printf("Roof type will be: %s\n", t.roofType)
}
