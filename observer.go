package observer

type Observer interface {
	OnNotify(data interface{})
}

type Func func(data interface{})
func (f Func) OnNotify(data interface{}) {
	f(data)
}
