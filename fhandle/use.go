package fhandle

import (
	"DIMSProxy/algo"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
)

func Use(req protocol.AlgoUseReq) {
	var res protocol.AlgoUseRes
	resp, err := http.Get(req.ProductUrl)
	if err != nil {
		log.Logger.Errorf("Send Get Request To Url Error:%s", err.Error())
		res.Code = protocol.FErrorCode
		res.Msg = err.Error()
		marshal, _ := json.Marshal(res)
		err := FrontConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("Write Front Use Response Error:%s", err.Error())
		}
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Errorf("Read Body Error:%s", err.Error())
	}
	executeFunc, err := algo.AlgoExecuteFunc(req.AlgoName, string(all), len(all))
	if err != nil {
		log.Logger.Errorf("AlgoExecuteFunc Error:%s", err.Error())
		return
	}
	res.Code = protocol.FSuccessCode
	res.Msg = protocol.FSuccessMsg
	res.Data.HaveData = true
	res.Data.Data = executeFunc
	marshal, _ := json.Marshal(res)
	err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("Write Front Use Response Error:%s", err.Error())
		return
	}
}
