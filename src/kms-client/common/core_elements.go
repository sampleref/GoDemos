package common


var ResponseChannels = make(map[float64]chan Response)

// Error that can be filled in response
type Error struct {
	Code    int64
	Message string
	Data    string
}

// Response represents server response
type Response struct {
	Jsonrpc string
	Id      float64
	Method	string
	Result  map[string]string // should change if result has no several form
	Params  map[string]string
	Error   *Error
}

