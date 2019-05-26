package models

//ApiMessage api response
type ApiMessage struct {
	Code    int
	Success bool
	Message string
	Data    interface{}
}
