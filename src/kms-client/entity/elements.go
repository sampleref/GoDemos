package entity

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
	Result  map[string]string // should change if result has no several form
	Params  map[string]string
	Error   *Error
}
