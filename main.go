package main

import (
	"DIMSProxy/config"
	"DIMSProxy/file"
	"DIMSProxy/handle"
	"DIMSProxy/log"
	"DIMSProxy/manager"
	"DIMSProxy/model"
	"DIMSProxy/service"
	"DIMSProxy/util"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ebitengine/purego"
)

func init() {
	//dev
	util.ReadCopy()
	handle.Chc = make(chan bool, 1)
	model.CertCh = make(chan bool, 1)
	handle.PInfoCh = make(chan []byte, 1)
	config.NewConfig()
	log.NewLogger()
	file.NewMinioClient()
	model.OpenPInfo()
	model.OpenPLog()
	model.OpenCert()
}
func ctrlc() {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		<-c
		log.Logger.Info("Ctrl +C")
		if handle.WsConn == nil {
			os.Exit(200)
		} else {
			err := handle.WsConn.Close()
			if err != nil {
				log.Logger.Error(err)
				os.Exit(200)
			}
		}
		os.Exit(200)
	}

}
func invoke() {
	position := config.GetCaLibPosition()
	libc, err := util.OpenLibrary(position)
	if err != nil {
		log.Logger.Errorf("加载证书动态库错误:%s", err.Error())
		return
	}
	var caRun func()
	purego.RegisterLibFunc(&caRun, libc, "CaRunning")
	model.CertCh <- true
	caRun()
}
func run() {
	http.HandleFunc("/api/v1/user/login", service.Connect)
	haddr := config.Conf.Local.Host + ":" + strconv.Itoa(config.Conf.Local.Port)
	log.Logger.Infof("monitor  service run addr:%v", haddr)
	err := http.ListenAndServe(haddr, nil)
	if err != nil {
		panic(err)
	}
}
func main() {
	//fmt.Println(config.Conf.Local.CurrentDir)

	go invoke()
	go run()
	handle.WsStatusCh = make(chan bool, 1)
	handle.Connect()
	time.Sleep(time.Second * 3)
	go handle.LopWsStatus()
	go handle.Reader()
	go file.Add()
	go file.Clean()
	go manager.WatcherRun()
	go manager.UdpListen()
	//go proxy.Run()
	go ctrlc()
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	////go handle.WsSend()
	//batchLog, s, _ := model.FindBatchLog(94)
	//log.Logger.Infof("记录----------------------------：%v:%v", batchLog, s)
	select {}
}
