package request

type Airdrop struct {
	Address string `json:"address" binding:"required"` // 空投地址
}
