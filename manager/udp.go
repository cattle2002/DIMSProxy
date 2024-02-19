package manager

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"encoding/json"
	"net"
)

type FindReq struct {
	Cmd string `json:"Cmd"`
	Ip  string `json:"Ip"`
}
type FindRes struct {
	Cmd string `json:"Cmd"`
	Ip  string `json:"Ip"`
}

func UdpListen() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":5519")
	if err != nil {
		log.Logger.Errorf("resolve udp address error::%s", err.Error())
		return
	}
	serverConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Logger.Errorf("listen udp addr error:%s", err.Error())
		return
	}
	for {
		var buff [512]byte
		n, rAddr, err := serverConn.ReadFromUDP(buff[0:])
		if err != nil {
			log.Logger.Errorf("Read From Udp Error:%s", err.Error())
			continue
		}

		log.Logger.Infof("remote udp addr:%v:%v", rAddr.IP, rAddr.Port)
		log.Logger.Tracef("read udp msg:%s", string(buff[:n]))

		var res FindRes

		var req FindReq
		err = json.Unmarshal(buff[:n], &req)
		if err != nil {
			log.Logger.Errorf("josn marshal msg  error:%s", err.Error())
			log.Logger.Tracef("read msg:%s", string(buff[:n]))
			res.Cmd = "FindRet"
			res.Ip = "protocol feild not support"
			b, _ := json.Marshal(res)
			log.Logger.Tracef("write  find Ret  to  manaager:%s", b)
			_, err = serverConn.WriteToUDP(b, rAddr)
			if err != nil {
				log.Logger.Errorf("write errir:%s", err.Error())
				return
			}
		}
		//res.Ip = config.Conf.Local.EthHost + ":" + strconv.Itoa(config.Conf.Local.ManagerServicePort) + "/api/v1/websocket/1"
		res.Ip = config.Conf.Local.EthHost
		res.Cmd = "FindRet"
		marshal, _ := json.Marshal(res)
		_, err = serverConn.WriteToUDP(marshal, rAddr)
		if err != nil {
			log.Logger.Errorf("write error:%s", err.Error())
			return
		}
		log.Logger.Tracef("Find Ret :%s", marshal)
		continue
	}
}
