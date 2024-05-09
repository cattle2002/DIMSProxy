package protocol

const (
	Keep = "Keep"

	FSuccessCode    = 200
	FErrorCode      = 400
	CoreSuccessCode = 0
	CoreErrorCode   = -400
	Login           = "Login"
	LoginRet        = "LoginRet"
	Monitor         = "Monitor"
	MonitorRet      = "MonitorRet"
	MonitorData     = "MonitorData"
	MonitorDataRet  = "MonitorDataRet"
	FSuccessMsg     = "success"
	Calc            = "Calc"
	CalcRet         = "CalcRet"
	OStatus         = "healthy"
	Stop            = "stop"
	Renew           = "renew"
	Logg            = "log"
	Time            = "times"
	ProductNotFound = "数据产品记录未找到"
	Limit           = "Limit"
	LimitRet        = "LimitRet"
	LocalCanUse     = "LocalCanUseSyncReq"
	LocalCanUseRet  = "LocalCanUseSyncRet"
	OffEncrypt      = "OffEncrypt"
	OffEncryptRet   = "OffEncryptRet"
	AlgoList        = "AlgoList"
	AlgoRegister    = "AlgoRegister"
	HttpAlgoListRet = "AlgoListRet"
	Use             = "Use"
	UseRet          = "UseRet"
)
