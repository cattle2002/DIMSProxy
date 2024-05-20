package fhandle

import (
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

func Use(req protocol.AlgoUseReq) {
	var res protocol.AlgoUseRes
	_, err := http.Get(req.ProductUrl)
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
	//all, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Logger.Errorf("Read Body Error:%s", err.Error())
	//}
	//executeFunc, err := algo.AlgoExecuteFunc(req.AlgoName, string(all), len(all))
	//if err != nil {
	//	log.Logger.Errorf("AlgoExecuteFunc Error:%s", err.Error())
	//	return
	//}
	//create, _ := os.Create("订单数据.json")
	//defer create.Close()
	//io.Copy(create, bytes.NewReader(all))

	//result, err := CalcResult(all)
	file, _ := os.ReadFile("out.json")
	//log.Logger.Errorf("use error:%s", err.Error())
	//create, _ := os.Create("use.txt")
	//reader := strings.NewReader(executeFunc)
	//io.Copy(create, reader)
	//header := "==========================此数据已经被算法处理==========================\n"
	//split := strings.Split(executeFunc, header)

	res.Cmd = protocol.UseRet
	res.Code = protocol.FSuccessCode
	res.Msg = protocol.FSuccessMsg
	res.Data.HaveData = true
	//res.Data.Data = result
	res.Data.Data = string(file)
	res.Data.Type = "json"
	marshal, _ := json.Marshal(res)
	err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("Write Front Use Response Error:%s", err.Error())
		return
	}
}
