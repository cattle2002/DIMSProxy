package service

import (
	"DIMSProxy/fhandle"
	"DIMSProxy/log"
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有的跨域请求，实际应用中需要根据需求配置
		return true
	},
} // use default options

func Connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Errorf("established fron websocket error:%s", err.Error())
		return
	}
	fhandle.FrontConn = c
	loopReader()
	defer fhandle.FrontConn.Close()
}
func loopReader() {
	for {
		log.Logger.Trace("reading  front websocket msg...")
		_, p, err := fhandle.FrontConn.ReadMessage()
		if err != nil {
			log.Logger.Errorf("reading  front websocket msg error:%s", err.Error())
			break
		}
		log.Logger.Tracef("read front msg:%s", string(p))
		cmd, err := packet.ExtractCmdValue(string(p))
		if err != nil {
			log.Logger.Errorf("read cmd error:%s", err.Error())
			continue
		}
		if cmd == protocol.Login { /*前端发送的登录请求*/
			var loginReq protocol.HttpLoginReq
			err := json.Unmarshal(p, &loginReq)
			if err != nil {
				log.Logger.Errorf("unmarshal front req error:%s", err.Error())
				continue
			}
			fhandle.Login(loginReq)
		}
		if cmd == protocol.Calc {
			var calcReq protocol.HttpCalcRequest
			err := json.Unmarshal(p, &calcReq)
			if err != nil {
				log.Logger.Errorf("unmarshal front req error:%s", err.Error())
				continue
			}
			fhandle.Calc(calcReq)
		}
		if cmd == protocol.OffEncrypt {
			var offencReq protocol.HttpOfflineEncryptReq
			err := json.Unmarshal(p, &offencReq)
			if err != nil {
				log.Logger.Errorf("unmarshal front req error:%s", err.Error())
				continue
			}
			fhandle.Encrypt(offencReq)
		}
		if cmd == protocol.AlgoList {
			fhandle.AlgoListX()
		}
		if cmd == protocol.AlgoRegister {
			fhandle.AlgoRegister()
		}
		if cmd == protocol.Use {
			var req protocol.AlgoUseReq
			err := json.Unmarshal(p, &req)
			if err != nil {
				log.Logger.Errorf("Json Unmarshal Error:%s", err.Error())
				continue
			}
			fhandle.Use(req)
		}
	}
}
