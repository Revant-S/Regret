package simdomain

// Evaluator owns the pool , the http request client
type Evaluator struct {
	MessagePool *pool
	Score       float64
}

type EvaluatorInterface interface {
	EvaluateMessageScore(messageId *string, carrier *CarrierResponse) float64
	EvaluateTotalScore() float64
}
