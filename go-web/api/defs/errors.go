package defs




type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,
			Error:Err{"request body is not correct","001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC:401,
			Error:Err{"user auth failed","002"}}
)
