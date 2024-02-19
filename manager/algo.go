package manager

import (
	"DIMSProxy/algo"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
)

func AlgoRegisterHandle(p []byte) {
	log.Logger.Infof("Manager Read AlgoRegister Request Msg:%s", string(p))
	var req protocol.AlgoRegisterReq
	err := json.Unmarshal(p, &req)
	if err != nil {
		AlgoRegisterResponseError("protocol error")
		return
	}
	registerFunc, err := algo.AlgoRegisterFunc(req.AlgoJson)
	if err != nil {
		AlgoRegisterResponseError(registerFunc)
		return
	}
	AlgoRegisterResponseSuccess(registerFunc)
}
func AlgoGetAllHandle(p []byte) {
	log.Logger.Infof("Manager Read AlgoGetAll Msg:%s", string(p))
	var req protocol.AlgoGetAllReq
	err := json.Unmarshal(p, &req)
	if err != nil {
		AlgoGetAllResponseError(err.Error())
		return
	}
	listFunc, err := algo.AlgoGetListFunc()
	if err != nil {
		AlgoGetAllResponseError(err.Error())
		return
	}
	AlgoGetAllResponseSuccess(listFunc)
}
func AlgoDeleteHandle(p []byte) {
	log.Logger.Infof("Manager Read AlgoDelete Msg:%s", string(p))
	var req protocol.AlgoDeleteReq
	err := json.Unmarshal(p, &req)
	if err != nil {
		AlgoDeleteResponseError(err.Error())
		return
	}
	deleteFunc, err := algo.AlgoDeleteFunc(req.AlgoName)
	if err != nil {
		AlgoDeleteResponseError(err.Error())
		return
	}
	AlgoDeleteResponseSuccess(deleteFunc)
}
