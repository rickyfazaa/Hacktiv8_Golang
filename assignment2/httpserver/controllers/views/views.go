package views

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

const (
	BAD_REQUEST           = "BAD_REQUEST"
	CREATED               = "CREATED"
	OK                    = "OK"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
)

func ErrorResp(status int, message string, error error) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Error:   error.Error(),
	}
}

func SuccessResp(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}
