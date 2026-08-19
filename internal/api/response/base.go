package response

type BaseResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](message string, data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func SuccessWithoutData(message string) *BaseResponse[struct{}] {
	return &BaseResponse[struct{}]{
		Success: true,
		Message: message,
	}
}

func Error(message string) *BaseResponse[struct{}] {
	return &BaseResponse[struct{}]{
		Success: false,
		Message: message,
	}
}
