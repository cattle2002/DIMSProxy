package manager

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var ManagerConn *websocket.Conn

var upgrader = websocket.Upgrader{} // use default options
func pkt(jsonStr []byte) (string, error) {
	mp := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &mp)
	if err != nil {
		log.Logger.Errorf("反序列化失败:%s", err.Error())
		return "", err
	}
	return mp["Cmd"].(string), nil
}
func Watch(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Errorf("upgrade Error:%s", err)
		return
	}
	log.Logger.Tracef("管理端连接:%s", c.RemoteAddr().String())
	ManagerConn = c
	handle()
}
func WatcherRun() {
	//todo 初始化库的位置
	//NewAlgoLib()
	http.HandleFunc("/api/v1/websocket/1", Watch)
	addr := config.Conf.Local.EthHost + ":" + strconv.Itoa(config.Conf.Local.ManagerServicePort)
	log.Logger.Infof("anager service run addr:%s/api/v1/websocket/1", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Logger.Errorf("manager service run failed:%s", err.Error())
		return
	}
}
