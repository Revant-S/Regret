package domain

type MessageResult bool

const TickHeaderString string = "X-TICK"

const (
	MessageResultSuccess MessageResult = true
	MessageResultFailure MessageResult = false
)

type Message struct {
	Id      string
	Text    string
	OutCome MessageResult
}
