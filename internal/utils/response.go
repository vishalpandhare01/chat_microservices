package utils

type ErrorResponse struct {
	Code    int
	Message string
}

type SuccessResponse struct {
	Code    int
	Message string
	Data    interface{}
}

type SuccessListResponse struct {
	Total   int
	Perpage int
	Page    int
	Data    interface{}
}
