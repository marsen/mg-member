package service

type AService struct {
	note string
}

type Caller interface {
	Call(i AInput) (o AOutput)
}

func NewAService() *AService {
	return &AService{note: "test"}
}

func (s AService) Call(i AInput) (o AOutput) {
	return AOutput{
		Result: "Sandbox",
	}
}

type AInput struct {
	Raw interface{} `json:"raw"`
}

type AOutput struct {
	Result string `json:"result"`
}
