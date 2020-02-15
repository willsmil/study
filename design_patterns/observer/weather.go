package observer

type Weather struct {
	Observers []Observer
	Temp      float64
	Humidity  float64
	Pressure  float64
}

func (w *Weather) SetWeather(temp, humidity, pressure float64) {
	w.Temp = temp
	w.Pressure = pressure
	w.Humidity = humidity
	w.WeatherChanged()
}
func (w *Weather) WeatherChanged() {
	w.NotifyObservers()
}

func (w *Weather) RegisterObserver(o Observer) {
	w.Observers = append(w.Observers, o)
}
func (w *Weather) RemoverObserver(o Observer) {
	for i, observer := range w.Observers {
		if o == observer {
			w.Observers = append(w.Observers[:i], w.Observers[i+1:]...)
		}
	}
}
func (w *Weather) NotifyObservers() {
	for _, o := range w.Observers {
		o.Update(w.Temp, w.Humidity, w.Pressure)
	}
}
