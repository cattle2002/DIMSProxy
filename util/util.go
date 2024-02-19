package util

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
	"time"
)

func FindFile(fn string) bool {
	if runtime.GOOS == "windows" {
		fp := config.Conf.Local.CurrentDir + "\\" + fn
		_, err := os.Stat(fp)
		if err != nil {
			return false
		} else {
			return true
		}
	} else {
		fp := config.Conf.Local.CurrentDir + "/" + fn
		_, err := os.Stat(fp)
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

// 返回文件的绝对路径
func returnCFileAbs(fn string) string {
	if runtime.GOOS == "windows" {
		fp := config.Conf.Local.CurrentDir + "\\" + fn
		return fp
	} else {
		fp := config.Conf.Local.CurrentDir + "/" + fn
		return fp
	}
}

//	func ReturnLFileAbs(fn string) string {
//		if runtime.GOOS == "windows" {
//			fp := config.Conf.Local.CurrentDir + "\\" + "lib\\" + fn
//			return fp
//		} else {
//			fp := config.Conf.Local.CurrentDir + "/" + "lib/" + fn
//			return fp
//		}
//	}
func ReturnLFileAbs(fn string) string {
	if runtime.GOOS == "windows" {
		fp := config.Conf.Local.CurrentDir + "\\" + fn
		return fp
	} else {
		fp := config.Conf.Local.CurrentDir + "/" + fn
		return fp
	}
}
func libcert() (string, string, error) {
	pk, err := os.ReadFile(ReturnLFileAbs(config.Conf.KeyPair.PublicKeyPath))
	if err != nil {
		return "", "", err
	}
	sk, err := os.ReadFile(ReturnLFileAbs(config.Conf.KeyPair.PrivateKeyPath))
	if err != nil {
		return "", "", err
	}
	return string(pk), string(sk), nil
}
func GetCurrentMillSecond(productTime int64) bool {
	ms := time.Now().Unix() * 1000
	if ms >= productTime {
		return true
	}
	return false
}
func CopyCertNew() {

}
func CopyCert() error {
	log.Logger.Trace("正在copy公私钥")
	_, _, err := libcert()
	if err != nil {
		return err
	}
	return nil
	//pw, err := os.Create(returnCFileAbs(config.Conf.KeyPair.PublicKeyPath))
	//if err != nil {
	//	return err
	//}
	//sw, err := os.Create(returnCFileAbs(config.Conf.KeyPair.PrivateKeyPath))
	//if err != nil {
	//	return err
	//}
	//pr := strings.NewReader(pk)
	//sr := strings.NewReader(sk)
	//_, err = io.Copy(pw, pr)
	//if err != nil {
	//	return err
	//}
	//_, err = io.Copy(sw, sr)
	//if err != nil {
	//	return err
	//}
	//return nil
}

func Md5(data []byte) string {
	s := md5.Sum(data)
	toString := hex.EncodeToString(s[:])
	return toString
}
func DateFormat(timestamp int64) string {
	tm := time.UnixMilli(timestamp)
	return tm.Format("2006-01-02 15:04:05")
}
