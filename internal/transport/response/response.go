package response

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

const (
	StatusOk  = "Ok"
	StatusErr = "Error"
)

func Ok() Response {
	return Response{
		Status: StatusOk,
	}
}

func Err(msg string) Response {
	return Response{
		Status:  StatusErr,
		Message: msg,
	}
}
