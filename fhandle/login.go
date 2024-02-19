package fhandle

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"encoding/json"
	"github.com/gorilla/websocket"
)

var FrontConn *websocket.Conn

func Login(req protocol.HttpLoginReq) {
	var res protocol.HttpLoginRes
	if config.Conf.Local.User == req.User.Username && config.Conf.Local.Password == req.User.Password {
		res.Cmd = protocol.LoginRet
		res.Code = protocol.FSuccessCode
		res.Msg = protocol.FSuccessMsg
		marshal, err := json.Marshal(res)
		if err != nil {
			log.Logger.Errorf("marshal login msg error:%s", err.Error())
			return
		}
		//todo copy lib下的公私钥
		err = util.CopyCert()
		if err != nil {
			log.Logger.Errorf("copy ca to monitor error:%s", err.Error())
			return
		}
		err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("write msg to front  error:%s", err.Error())
			return
		}
		return
	} else {
		res.Cmd = protocol.LoginRet
		res.Code = protocol.FErrorCode
		res.Msg = "用户名和密码不匹配"
		marshal, err := json.Marshal(res)
		if err != nil {
			log.Logger.Errorf("marshal login msg error:%s", err.Error())
			return
		}
		err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("write msg to front  error:%s", err.Error())
			return
		}
	}
}
