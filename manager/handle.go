package manager

import "DIMSProxy/log"

func handle() {
	for {
		log.Logger.Trace("reading manager msg...")
		_, p, err := ManagerConn.ReadMessage()
		if err != nil {
			log.Logger.Errorf("reading manager msg error:%s", err.Error())
			return
		}
		log.Logger.Tracef("read manager msg:%s", string(p))
		cmd, err := pkt(p)
		if err != nil {
			log.Logger.Errorf("Get Protocol Cmd Field Error:%s", err.Error())
			return
		}
		if cmd == string(AlgoRegister) {
			AlgoRegisterHandle(p)
		}
		if cmd == string(AlgoGetAll) {
			AlgoGetAllHandle(p)
		}
		if cmd == string(AlgoDelete) {
			AlgoDeleteHandle(p)
		}
		if cmd == string(CertInput) {
			CertInputResponseSuccess()
		}
		if cmd == string(CertRemake) {
			CertRemakeResponseSuccess()
		}
		if cmd == string(CertSync) {
			CertSyncResponseSuccess()
		}
		if cmd == string(CertShow) {
			CertShowHandle()
		}
		if cmd == string(CertOwner) {
			CertOwnerHandle()
		}
	}
}
