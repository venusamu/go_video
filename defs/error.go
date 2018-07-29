package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorRespons struct {
	HttpStatusCode int
	Error          Err
}

var (
	BodyRequestParseFailed = ErrorRespons{
		HttpStatusCode: 400,
		Error: Err{
			Error:     "parse body failed",
			ErrorCode: "001",
		},
	}

	NotAuthUser = ErrorRespons{
		HttpStatusCode: 401,
		Error: Err{
			Error:     "User not exist",
			ErrorCode: "002",
		},
	}
)
