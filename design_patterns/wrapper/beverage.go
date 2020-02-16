package wrapper

type Beverage interface {
	Cost() float64
	GetDescription() string
}

type Description struct {
	Description string
}
