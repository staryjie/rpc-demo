package service

const (
	SERVICE_NAME = "HelloService"
)

type HelloService interface {
	Hello(request string, response *string) error
	Calc(request *CalcRequest, response *int) error
}

type CalcRequest struct {
	A int `json: "a"`
	B int `json: "b"`
}
