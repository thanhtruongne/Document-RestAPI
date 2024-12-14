package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

type SimpleResponse struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {

	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *SimpleResponse {
	return &SimpleResponse{Data: data}
}
