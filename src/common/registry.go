package common

// HandlerResult indicates what the handler did
type HandlerResult struct {
	handled bool
}

// Handled means the handler took care of execution (called ExecReal or handled it entirely)
var Handled = HandlerResult{handled: true}

// PassThru means pass through to the real command with original args
var PassThru = HandlerResult{handled: false}

// IsHandled returns true if the handler handled the command
func (r HandlerResult) IsHandled() bool {
	return r.handled
}

var Handlers = map[string]func() HandlerResult{}

func Register(name string, handler func() HandlerResult) {
	Handlers[name] = handler
}
