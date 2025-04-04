package dto

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Payload  T  `json:"payload"`
}