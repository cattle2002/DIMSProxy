package manager

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func AlgoRegisterResponseError(data string) {
	var res protocol.AlgoRegisterRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Msg = data
	res.Cmd = string(AlgoRegisterRet)
	res.Data = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager :%s", string(marshal))
}
func AlgoRegisterResponseSuccess(data string) {
	var res protocol.AlgoRegisterRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Msg = protocol.FSuccessMsg
	res.Cmd = string(AlgoRegisterRet)
	res.Data = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager :%s", string(marshal))
}

func AlgoGetAllResponseError(data string) {
	var res protocol.AlgoGetAllRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Msg = data
	res.Cmd = string(AlgoGetAllRet)
	res.Data = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager :%s", string(marshal))
}
func AlgoGetAllResponseSuccess(data string) {
	var res protocol.AlgoGetAllRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(AlgoGetAllRet)
	res.Msg = protocol.FSuccessMsg
	res.Data = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func AlgoDeleteResponseError(data string) {
	var res protocol.AlgoDeleteRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Cmd = string(AlgoDeleteRet)
	res.Msg = data
	res.Data = ""
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func AlgoDeleteResponseSuccess(data string) {
	var res protocol.AlgoDeleteRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(AlgoDeleteRet)
	res.Msg = protocol.FSuccessMsg
	res.Data = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func CertInputResponseError(data string) {
	var res protocol.CertInputRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Cmd = string(CertInputRet)
	res.Msg = data
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func CertInputResponseSuccess() {
	var res protocol.CertInputRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(CertInputRet)
	res.Msg = protocol.FSuccessMsg
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func CertRemakeResponseSuccess() {
	var res protocol.CertRemakeRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(CertRemakeRet)
	res.Msg = protocol.FSuccessMsg
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
func CertSyncResponseSuccess() {
	var res protocol.CertSyncRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(CertSyncRet)
	res.Msg = protocol.FSuccessMsg
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}

func CertShowResponseError(errMsg string) {
	var res protocol.CertShowRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Cmd = string(CertShowRet)
	res.Msg = errMsg
	res.Msg = protocol.FSuccessMsg
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}

func CertOwnerResponseError(errMsg string) {
	var res protocol.CertOwnerRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FErrorCode
	res.Cmd = string(CertOwnerRet)
	res.Msg = errMsg
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager :%s", string(marshal))
}
func CertOwnerResponseSuccess(pk string, sk string) {
	var res protocol.CertOwnerRes
	res.IpAddr = config.Conf.Local.EthHost
	res.Code = protocol.FSuccessCode
	res.Cmd = string(CertOwnerRet)
	res.Msg = protocol.FSuccessMsg
	res.PublicKey = pk
	res.PrivateKey = sk
	marshal, _ := json.Marshal(res)
	err := ManagerConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return msg to manager error:%s", err.Error())
		return
	}
	log.Logger.Tracef("return msg to manager:%s", string(marshal))
}
