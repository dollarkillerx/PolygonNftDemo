package enum

type BasicInformation string

func (b BasicInformation) String() string {
	return string(b)
}

const (
	RequestID    BasicInformation = "request_id"
	AuthModel    BasicInformation = "authModel"
	AuthMqlModel BasicInformation = "authMqlModel"
)
