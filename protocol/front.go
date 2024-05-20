package protocol

type User struct {
	Username string `json:"name"`
	Password string `json:"password"`
}
type HttpLoginReq struct {
	Cmd  string `json:"Cmd"`
	User User   `json:"user"`
}
type HttpLoginRes struct {
	Cmd  string `json:"Cmd"`
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
}

type HttpCalcResponsePayload struct {
	ID       int64  `json:"ID"`
	HaveData bool   `json:"HaveData"`
	Data     []byte `json:"Data"`
	Url      string `json:"Url"`
	//CalcResult string `json:"CalcResult"`
}
type HttpCalcResponse struct {
	Cmd      string                  `json:"Cmd"`
	RetCode  int                     `json:"RetCode"`
	ErrorMsg string                  `json:"Msg"`
	Payload  HttpCalcResponsePayload `json:"Payload"`
}
type HttpCalcRequestPayload struct {
	ID                                    int64  `json:"ID"`
	ProductID                             int64  `json:"ProductID"`
	ProductName                           string `json:"ProductName"`
	Buyer                                 string `json:"Buyer"`
	Seller                                string `json:"Seller"`
	CertificateAlgorithmTypeAlgorithmType string `json:"CertificateTypeAlgorithmType"`
	SymmetricKeyAlgorithmType             string `json:"SymmetricKeyAlgorithmType"`
	HaveData                              bool   `json:"HaveData"`
	ProductType                           string `json:"ProductType"`
	ProductData                           string `json:"ProductData"`
	ProductUrl                            string `json:"ProductPosition"`
	CipherSymmetricKey                    string `json:"CipherSymmetricKey"`
	BuyerCaTimeStamp                      int64  `json:"BuyerCaTimeStamp"`
	SellerCaTimeStamp                     int    `json:"SellerCaTimeStamp"`
	SellerKey                             string `json:"SellerKey"`
	NeedConfirmAndGrantPermission         bool   `json:"NeedConfirmAndGrantPermission"`
	NeedConfirmPermission                 bool   `json:"NeedConfirmPermission"`
	NeedGrantPermission                   bool   `json:"NeedGrantPermission"`
	NeedDecrypt                           bool   `json:"NeedDecrypt"`
}
type HttpCalcRequest struct {
	Cmd     string                 `json:"Cmd"`
	Payload HttpCalcRequestPayload `json:"Payload"`
}
type HttpOfflineEncryptReq struct {
	FileUrl    string `json:"FileUrl"`
	Encryption string `json:"Encryption"`
	Password   string `json:"Password"`
}
type HttpOfflineEncryptRes struct {
	Cmd  string                    `json:"Cmd"`
	Code int                       `json:"Code"`
	Msg  string                    `json:"Msg"`
	Data HttpOfflineEncryptResData `json:"Data"`
}
type HttpOfflineEncryptResData struct {
	Url        string `json:"Url"`
	AlgoType   string `json:"AlgoType"`
	ConfirmStr string `json:"ConfirmStr"`
}
type AlgoListRes struct {
	Cmd  string `json:"Cmd"`
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
}
