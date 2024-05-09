package main

import (
	"DIMSProxy/util"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ebitengine/purego"
	"time"
)

type AlgoModel struct {
	Name         string    `gorm:"primaryKey" json:"Name"`
	Type         string    `json:"Type"`
	FilePath     string    `json:"FilePath"`
	StartupCmd   string    `json:"StartupCmd"`
	ExeInputEOF  string    `json:"ExeInputEOF"`
	MaxExecTime  int       `json:"MaxExecTime"`
	InputExample string    `json:"InputExample"`
	AlgoFuncName string    `json:"AlgoFuncName"`
	Parameters   string    `json:"Parameters"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

type AlgoLogModel struct {
	ID         uint      `gorm:"primaryKey" json:"ID"`
	AlgoName   string    `json:"AlgoName"`
	AlgoType   string    `json:"AlgoType"`
	ExecTime   time.Time `json:"ExecTime"`
	ExecResult string    `json:"ExecResult"`
	ExecCost   int       `json:"ExecCost"`
}

const ss = ""

func AlgoRegisterFunc(algoJson string) (string, error) {
	//log.Logger.Trace("begin invoke AlgoRegister API")
	//startTime := time.Now()
	algoPos := "D:\\workdir\\DIMSProxy\\algo\\libalgo.dll"
	libc, err := util.OpenLibrary(algoPos)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := util.CloseLibrary(libc); cerr != nil {
			panic(err)
		}
	}()
	var AlgoRegisterName func(string) string
	purego.RegisterLibFunc(&AlgoRegisterName, libc, "RegisterAlgo")
	rs := AlgoRegisterName(algoJson)
	//endTime := time.Now()
	//duration := endTime.Sub(startTime)
	//seconds := int(duration.Seconds())
	return rs, nil

}
func main() {
	model := AlgoModel{
		Name:         "算法名称",
		Type:         "EXE",
		FilePath:     "D:\\workdir\\DIMSProxy\\algo\\algo3.exe",
		StartupCmd:   "",
		ExeInputEOF:  "",
		MaxExecTime:  20,
		InputExample: "",
		AlgoFuncName: "",
		Parameters:   "",
		//CreatedAt:    time.Time{},
	}
	marshal, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
	//registerFunc, err := AlgoRegisterFunc(string(marshal))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(registerFunc)
	//{
	//	// 算法名称(唯一标识)
	//	"Name": "",
	//// 算法类型(EXE,LIB,API)
	//	"Type": "",
	//// 算法文件路径(可执行文件路径或者动态库路径)
	//	"FilePath": "",
	//// 可选的算法程序启动命令
	//// 例如调用./main.py，StartupCmd的值为"python",FilePath的值为"./main.py"
	//	"StartupCmd": "",
	//// 可选的输入终止字符串，用于算法程序判断输入是否结束(例如 #======EOF======#)
	//	"ExeInputEOF": "",
	//// 算法程序最大执行时间，单位秒。超过该时间算法程序将被强制终止并返回超时
	//	"MaxExecTime": 0,
	//// 输入示例(仅用于展示，无实际用途)
	//	"InputExample": "",
	//// 算法函数名称(仅用于动态库)
	//	"AlgoFuncName": "",
	//// 算法参数
	//// 调用算法程序时，会将该参数传递给算法程序(例如"-t 30 -m 1024")
	//// 调用动态库时，会将该参数传递给算法函数(可以是任意字符串，不做规定)
	//	"Parameters": "",
	//// 算法的创建时间，注册算法时不需要该字段(否则报错)
	//	"CreatedAt": "2023-11-29T19:48:32.3042678+08:00"
	//}
	s := "-----BEGIN Public key-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArTBCFNMP6/suAw1jOcsCZ7HNYWXbWgv8xYmaEIq6wybqLenx8qcb7eLRURJ95yLmy6FI/yK2riunY5Rlv5hb3cPDcuIBqHauYPdniRUvaO6/iyRJ0DXKvykkS7+nEft26DDDQ8WCPZaRjCDAEkTdNZMC5tDNpTfdJQob7xbRMpRHAnscxcyyIJbdeL2A8IXyykFa7AIyRuzZh4VlH1g8q5Va/txzsjwTYFxCU7lzFFo6qdgl//Iqs+fzMfPs1JOBiv8v92Zg4T0hKIH7BzLyP/gNb8Yp3XYuPu7D2D9QMw90yN8iGJCQQtaARsdM0RfOgm6olryp1xNS02eTFz32ZwIDAQAB\n-----END Public key-----"
	sum := md5.Sum([]byte(s))
	toString := hex.EncodeToString(sum[:])
	fmt.Println(toString)
}
