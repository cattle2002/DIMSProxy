package protocol

type LoginReqPayload struct {
	ID       int64  `json:"ID"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Time     int64  `json:"Time"`
}
type KeepLiveReq struct {
	Cmd     string             `json:"Cmd"`
	Program string             `json:"Program"`
	Payload KeepLiveReqPayload `json:"Payload"`
}
type KeepLiveReqPayload struct {
	ID        int64  `json:"ID"`
	User      string `json:"User"`
	LoginCode int    `json:"LoginCode"`
}

// type KeepLiveRes struct {
// }
type LoginReq struct {
	Cmd     string          `json:"Cmd"`
	Program string          `json:"Program"`
	Payload LoginReqPayload `json:"Payload"`
}
type LoginResPayload struct {
	ID        int64 `json:"ID"`
	TimeStamp int64 `json:"TimeStamp"`
}
type LoginRes struct {
	Cmd      string          `json:"Cmd"`
	Program  string          `json:"Program"`
	RetCode  int             `json:"RetCode"`
	ErrorMsg string          `json:"ErrorMsg"`
	Payload  LoginResPayload `json:"Payload"`
}
type MonitorProductReqPayload struct {
	ID        int64  `json:"ID"`
	Buyer     string `json:"Buyer"`
	ProductID int64  `json:"ProductID"`
	Demand    string `json:"Demand"`
}
type MonitorProductReq struct {
	Cmd     string                   `json:"Cmd"`
	Payload MonitorProductReqPayload `json:"Payload"`
}

type MonitorProductRes struct {
	Command  string                   `json:"Cmd"`
	Program  string                   `json:"Program"`
	RetCode  int                      `json:"RetCode"`
	ErrorMsg string                   `json:"ErrorMsg"`
	Payload  MonitorProductResPayload `json:"Payload"`
}
type MonitorProductResPayload struct {
	ID        int64  `json:"ID"`
	Buyer     string `json:"Buyer"`
	Seller    string `json:"Seller"`
	ProductID int64  `json:"ProductID"`
	Demand    string `json:"Demand"`
	Log       string `json:"Log"`
}
type LimitReqPayload struct {
	ID          int64  `json:"ID"`
	Buyer       string `json:"Buyer"`
	ProductName string `json:"ProductName"`
	ProductID   int64  `json:"ProductID"`
}
type LimitReq struct {
	Cmd     string          `json:"Cmd"`
	Program string          `json:"Program"`
	Payload LimitReqPayload `json:"Payload"`
}
type LimitResPayload struct {
	ID     int64          `json:"ID"`
	Way    int            `json:"Way"`
	Detail LimitResDetail `json:"Detail"`
}
type LimitResDetail struct {
	CanUseNumberLocal int64 `json:"CanUseNumberLocal"`
	CanUseTimeLocal   int64 `json:"CanUseTimeLocal"`
}
type LimitRes struct {
	Cmd      string          `json:"Cmd"`
	Program  string          `json:"Program"`
	RetCode  int             `json:"RetCode"`
	ErrorMsg string          `json:"ErrorMsg"`
	Payload  LimitResPayload `json:"Payload"`
}
type LocalCanUseSyncReqPayload struct {
	ID                int64  `json:"ID"`
	ProductID         int64  `json:"ProductID"`
	UserName          string `json:"UserName"`
	CanUseNumberLocal int64  `json:"CanUseNumberLocal"`
}
type LocalCanUseSyncReq struct {
	Cmd     string                    `json:"Cmd"`
	Program string                    `json:"Program"`
	Payload LocalCanUseSyncReqPayload `json:"Payload"`
}
type LocalCanUseSyncResPayload struct {
	ID int64 `json:"ID"`
}
type LocalCanUseSyncRes struct {
	Cmd      string                    `json:"Cmd"`
	Program  string                    `json:"Program"`
	RetCode  int                       `json:"RetCode"`
	ErrorMsg string                    `json:"ErrorMsg"`
	Payload  LocalCanUseSyncResPayload `json:"Payload"`
}
type AlgoUseReq struct {
	Cmd        string `json:"Cmd"`
	ProductUrl string `json:"ProductUrl"`
	AlgoName   string `json:"AlgoName"`
}
type AlgoUseResPayload struct {
	HaveData bool   `json:"HaveData"`
	Data     string `json:"Data"`
	Url      string `json:"Url"`
}
type AlgoUseRes struct {
	Cmd  string            `json:"Cmd"`
	Code int               `json:"Code"`
	Msg  string            `json:"Msg"`
	Data AlgoUseResPayload `json:"Data"`
}
