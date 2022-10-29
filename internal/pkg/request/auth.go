package request

type AuthJWT struct {
	Account string `json:"account"`
	Name    string `json:"name"`
}
