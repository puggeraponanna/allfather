package httputil

type Request struct{}

type Response[T any] struct {
	Data      T      `json:"data,omitempty"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
