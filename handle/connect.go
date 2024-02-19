package handle

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"context"
	"github.com/gorilla/websocket"
	"time"
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, config.Conf.PlatformUrl, nil)
	if err != nil {
		WsStatusCh <- false
		WsStatusBool = false
		return
	}

	log.Logger.Info("connect coreServer success")
	WsStatusBool = true
	WsStatusCh <- true
	WsConn = conn

	err = login()
	if err != nil {
		log.Logger.Errorf("login error:%s", err.Error())
		WsStatusBool = false
		WsStatusCh <- false
		return
	}
}
