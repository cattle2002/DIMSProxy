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

func KeepLive() {
	var kp protocol.KeepLiveReq
	kp.Cmd = protocol.Keep
	kp.Program = protocol.Monitor
	kp.Payload.ID = util.MsgID()
	kp.Payload.User = config.Conf.Local.User
	kp.Payload.LoginCode = RetCode
	marshal, _ := json.Marshal(kp)

	msg, err := packet.Packet(marshal)
	if err != nil {
		log.Logger.Errorf("Packet Error:%s", err.Error())
		return
	}

	err = WsConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Logger.Errorf("KeepLive Error:%s", err.Error())
		return
	}
}
func LopWsStatus() {
	go func() {
		for {
			<-time.After(time.Minute * 5)
			log.Logger.Tracef("keep liveing...")
			//go Connect()
			go KeepLive()
		}
	}()
	for {
		f := <-WsStatusCh
		if f == false {
			//log.Logger.Errorf("核心服务器掉线:%s", err.Error())
			log.Logger.Info("reconnect coreServer ...")
			go Connect()
		}
	}
}
