package simdomain

// Evaluator owns the pool
type Evaluator struct {
	MessagePool *Pool
	Score       float64
}

type EvaluatorInterface interface {
	EvaluateMessageScore(messageId *string, carrier *CarrierResponse) float64
	EvaluateTotalScore() float64
}
