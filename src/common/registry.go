package common

var Handlers = map[string]func(){}

func Register(name string, handler func()) {
	Handlers[name] = handler
}
