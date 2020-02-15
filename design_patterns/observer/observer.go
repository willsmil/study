package observer

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Display interface {
	Display()
}

type O1 struct {
}

func (o *O1) Update(temp, humidity, pressure float64) {
	println("temp:", temp, "humidity:", humidity, "pressure:", pressure)
}
func (o *O1) Display() {

}
