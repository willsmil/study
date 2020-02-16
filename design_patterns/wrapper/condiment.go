package wrapper

type CondimentDecorator interface {
	GetDescription() string
}

type Mocha struct {
	Beverage
}

func NewMocha(b Beverage) *Mocha {
	return &Mocha{
		Beverage: b,
	}
}

func (m *Mocha) Cost() float64 {
	return 1.0 + m.Beverage.Cost()
}
func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}

type Whip struct {
	Beverage
}

func NewWhip(b Beverage) *Whip {
	return &Whip{
		Beverage: b,
	}
}

func (w *Whip) Cost() float64 {
	return 1.0 + w.Beverage.Cost()
}
func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whip"
}
