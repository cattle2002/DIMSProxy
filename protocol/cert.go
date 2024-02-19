package protocol

type Cert2fReq struct {
	User       string `json:"User"`
	TimeStamp  int64  `json:"TimeStamp"`
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}
type Cert2fRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
}

//
