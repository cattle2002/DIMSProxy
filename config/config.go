package config

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type KeyPair struct {
	AutoConfig     bool   `json:"AutoConfig"`
	Algorithm      string `json:"Algorithm"`
	Bits           int    `json:"Bits"`
	PublicKeyPath  string `json:"PublicKeyPath"`
	PrivateKeyPath string `json:"PrivateKeyPath"`
}

type Local struct {
	Host               string `json:"Host"`               //主机IP地址
	Port               int    `json:"Port"`               //主机端口地址
	CertPort           int    `json:"CertPort"`           //证书服务的端口
	User               string `json:"User"`               //用户名
	Password           string `json:"Password"`           //用户密码
	CurrentDir         string `json:"CurrentDir"`         //当前程序的工作目录
	Version            string `json:"Version"`            //监管程序版本
	ManagerServicePort int    `json:"ManagerServicePort"` //管理服务的运行端口
	EthHost            string `json:"EthHost"`            //物理IP地址
	LoggerLevel        string `json:"LoggerLevel"`        //日志级别
	NoConsole          bool   `json:"NoConsole"`
}

type FileServer struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Port     int    `json:"Port"`
	Host     string `json:"Host"`
	RootDir  string `json:"RootDir"`
}

type Minio struct {
	LifeDay                   int    `json:"LifeDay"`
	EndPoint                  string `json:"EndPoint"`        //文件服务器minio的直连地址
	AccessKeyID               string `json:"AccessKeyID"`     //minio的用户名
	SecretAccessKey           string `json:"SecretAccessKey"` //minio的用户密码
	UseSSL                    bool   `json:"UseSSL"`          //minio是否启动ssl
	ProductBucket             string `json:"ProductBucket"`   //数据产品的上传桶名
	ProductUpload             string `json:"ProductUpload"`
	OfflineProductDataEncrypt string `json:"OfflineProductDataEncrypt"`
}
type Config struct {
	PlatformUrl string  `json:"PlatformUrl"`
	KeyPair     KeyPair `json:"KeyPair"`
	Local       Local   `json:"Local"`
	Minio       Minio   `json:"Minio"`
}

var Conf Config
var MonitorHttpServer string

func NewConfig() {
	UpdateConfigCwd()
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &Conf)
	if err != nil {
		panic(err)
	}
	ip := Conf.Local.EthHost
	port := Conf.Local.ManagerServicePort + 2
	sport := strconv.Itoa(port)
	MonitorHttpServer = fmt.Sprintf("http://%s:%s%s", ip, sport, "/api/v1/monitor")
}

//	func GetCaLibPosition() string {
//		if runtime.GOOS == "windows" {
//			return Conf.Local.CurrentDir + "\\lib\\libca.dll"
//		} else {
//			return Conf.Local.CurrentDir + "/lib/libca.so"
//		}
//	}
func GetCaLibPosition() string {
	if runtime.GOOS == "windows" {
		return Conf.Local.CurrentDir + "\\" + "libca.dll"
	} else {
		return Conf.Local.CurrentDir + "/" + "libca.so"
	}
}
func GetPrivateKeyPem() (string, error) {
	if runtime.GOOS == "windows" {
		privateFilePosition := Conf.Local.CurrentDir + "\\" + Conf.KeyPair.PrivateKeyPath
		file, err := os.ReadFile(privateFilePosition)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	} else {
		privateFilePosition := Conf.Local.CurrentDir + "/" + Conf.KeyPair.PrivateKeyPath
		file, err := os.ReadFile(privateFilePosition)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	}
}
func GetPublicKeyPem() (string, error) {
	if runtime.GOOS == "windows" {
		publicKeyPem := Conf.Local.CurrentDir + "\\" + Conf.KeyPair.PublicKeyPath
		file, err := os.ReadFile(publicKeyPem)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	} else {
		publicKeyPem := Conf.Local.CurrentDir + "/" + Conf.KeyPair.PublicKeyPath
		file, err := os.ReadFile(publicKeyPem)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	}
}
