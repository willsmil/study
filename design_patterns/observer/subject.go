package observer

type Subject interface {
	RegisterObserver(o Observer)
	RemoverObserver(o Observer)
	NotifyObservers()
} 