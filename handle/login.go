package handle

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

func login() error {
	var loginReq protocol.LoginReq
	loginReq.Cmd = protocol.Login
	loginReq.Program = protocol.Monitor
	loginReq.Payload.ID = util.MsgID()
	loginReq.Payload.Time = time.Now().UnixMilli()
	loginReq.Payload.User = config.Conf.Local.User
	loginReq.Payload.Password = config.Conf.Local.Password
	marshal, err := json.Marshal(loginReq)
	if err != nil {
		return err
	}
	msg, err := packet.Packet(marshal)
	if err != nil {
		return err
	}
	err = WsConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}
	log.Logger.Tracef("send  login msg:%s", string(msg))
	return nil
}

func handleLogin() {

}
