package wrapper

type Espresso struct {
	Description
}

func NewEspresso() Espresso {
	return Espresso{Description{Description: "Espresso"}}
}
func (e *Espresso) Cost() float64 {
	return 1.0
}
func (e *Espresso) GetDescription() string {
	return e.Description.Description
}

type HouseBlend struct {
	Description
}

func NewHouseBlend() Beverage {
	return &Espresso{Description{Description: "House Blend Coffee"}}
}
func (e *HouseBlend) Cost() float64 {
	return 1.0
}
func (e *HouseBlend) GetDescription() string {
	return e.Description.Description
}

type DarkRoast struct {
	Description
}

func NewDarkRoast() Beverage {
	return &Espresso{Description{Description: "Dark Roast Coffee"}}
}
func (e *DarkRoast) Cost() float64 {
	return 1.0
}
func (e *DarkRoast) GetDescription() string {
	return e.Description.Description
}

type Decaf struct {
	Description
}

func NewDecaf() Beverage {
	return &Espresso{Description{Description: "Decaf"}}
}
func (e *Decaf) Cost() float64 {
	return 1.0
}
func (e *Decaf) GetDescription() string {
	return e.Description.Description
}
