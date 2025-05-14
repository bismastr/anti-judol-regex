package handler

type Response struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		Status: status,
		Data:   data,
	}

}
