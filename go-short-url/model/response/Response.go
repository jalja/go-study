package response

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func Success(data interface{}) *Response {
	re := Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
	return &re
}
func Error(message string) *Response {
	re := Response{
		Code:    -1,
		Message: message,
	}
	return &re
}
