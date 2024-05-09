package algo

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/util"
	"runtime"
	"time"

	"github.com/ebitengine/purego"
)

type AlgoExportAPI string

const (
	AlgoExportAPIAlgoErrorPrefix AlgoExportAPI = " ALGORITHM_LIB_ERROR"
	AlgoExportAPIAlgoRegister    AlgoExportAPI = "RegisterAlgo"
	AlgoExportAPIAlgoUpdate      AlgoExportAPI = "UpdateAlgo"
	AlgoExportAPIAlgoGet         AlgoExportAPI = "GetAlgo"
	AlgoExportAPIAlgoList        AlgoExportAPI = "GetAlgoList"
	AlgoExportAPIAlgoDelete      AlgoExportAPI = "DeleteAlgo"
	AlgoExportAPIAlgoExecute     AlgoExportAPI = "ExecuteAlgo"
	AlgoExportAPIAlgosExecute    AlgoExportAPI = "ExecuteAlgos"
)

func GetAlgoLib() string {
	if runtime.GOOS == "windows" {
		fp := config.Conf.Local.CurrentDir + "\\algo\\" + "libalgo.dll"
		return fp
	} else {
		fp := config.Conf.Local.CurrentDir + "/algo/" + "libalgo.so"
		return fp
	}
}
func AlgoRegisterFunc(algoJson string) (string, error) {
	log.Logger.Trace("begin invoke AlgoRegister API")
	startTime := time.Now()
	algoPos := GetAlgoLib()
	libc, err := util.OpenLibrary(algoPos)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := util.CloseLibrary(libc); cerr != nil {
			log.Logger.Errorf("close algo lib error:%s", err.Error())
		}
	}()
	var AlgoRegisterName func(string) string
	purego.RegisterLibFunc(&AlgoRegisterName, libc, string(AlgoExportAPIAlgoRegister))
	rs := AlgoRegisterName(algoJson)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := int(duration.Seconds())
	log.Logger.Tracef("invoke AlgoRegister spend time:%d 秒", seconds)
	return rs, nil
}
func AlgoGetListFunc() (string, error) {
	log.Logger.Trace("begin invoke AlgoGetList API")
	startTime := time.Now()
	//algoPos := config.GetAlgoLibPosition()
	algoLib := GetAlgoLib()
	//log.Logger.Tracef("算法库位置:%s", algoLib)
	libc, err := util.OpenLibrary(algoLib)
	if err != nil {
		return "", err
	}
	// defer util.CloseLibrary(libc)
	defer func() {
		if err := util.CloseLibrary(libc); err != nil {
			log.Logger.Errorf("close algo lib error:%s", err.Error())
		}
	}()
	var AlgoGetListFuncName func() string
	purego.RegisterLibFunc(&AlgoGetListFuncName, libc, string(AlgoExportAPIAlgoList))
	rs := AlgoGetListFuncName()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := int(duration.Seconds())
	log.Logger.Tracef("invoke AlgoGetList spend time:%d 秒", seconds)
	return rs, nil
}
func AlgoExecuteFunc(algoName string, input string, inputLen int) (string, error) {
	log.Logger.Trace("begin invoke AlgoExecute API")
	startTime := time.Now()
	algoLib := GetAlgoLib()
	libc, err := util.OpenLibrary(algoLib)
	if err != nil {
		return "", err
	}
	// defer util.CloseLibrary(libc)
	defer func() {
		if cerr := util.CloseLibrary(libc); cerr != nil {
			log.Logger.Errorf("close algo lib error:%s", err.Error())
		}
	}()
	var AlgoExecuteFuncName func(string, string, int) string
	purego.RegisterLibFunc(&AlgoExecuteFuncName, libc, string(AlgoExportAPIAlgoExecute))
	rs := AlgoExecuteFuncName(algoName, input, inputLen)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := int(duration.Seconds())
	log.Logger.Tracef("invoke AlgoExecute spend time:%d 秒", seconds)
	return rs, nil
}
func AlgoDeleteFunc(algoJson string) (string, error) {
	log.Logger.Trace("begin invoke AlgoDelete API")
	startTime := time.Now()
	algoPos := GetAlgoLib()
	libc, err := util.OpenLibrary(algoPos)
	if err != nil {
		return "", err
	}
	// defer util.CloseLibrary(libc)
	defer func() {
		if cerr := util.CloseLibrary(libc); cerr != nil {
			log.Logger.Errorf("close algo lib error:%s", err.Error())
		}
	}()
	var AlgoDeleteFuncName func(string) string
	purego.RegisterLibFunc(&AlgoDeleteFuncName, libc, string(AlgoExportAPIAlgoDelete))
	rs := AlgoDeleteFuncName(algoJson)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := int(duration.Seconds())
	log.Logger.Tracef("invoke AlgoDelete spend time :%d 秒", seconds)
	return rs, nil
}
