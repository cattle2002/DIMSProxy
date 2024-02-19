package util

type KeyPairM struct {
	AutoConfig     bool   `json:"AutoConfig"`
	Algorithm      string `json:"Algorithm"`
	Bits           int    `json:"Bits"`
	PublicKeyPath  string `json:"PublicKeyPath"`
	PrivateKeyPath string `json:"PrivateKeyPath"`
}

type LocalM struct {
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

type FileServerM struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Port     int    `json:"Port"`
	Host     string `json:"Host"`
	RootDir  string `json:"RootDir"`
}

type MinioM struct {
	LifeDay                   int    `json:"LifeDay"`
	EndPoint                  string `json:"EndPoint"`        //文件服务器minio的直连地址
	AccessKeyID               string `json:"AccessKeyID"`     //minio的用户名
	SecretAccessKey           string `json:"SecretAccessKey"` //minio的用户密码
	UseSSL                    bool   `json:"UseSSL"`          //minio是否启动ssl
	ProductBucket             string `json:"ProductBucket"`   //数据产品的上传桶名
	ProductUpload             string `json:"ProductUpload"`
	OfflineProductDataEncrypt string `json:"OfflineProductDataEncrypt"`
}
type ConfigM struct {
	PlatformUrl string   `json:"PlatformUrl"`
	KeyPair     KeyPairM `json:"KeyPair"`
	Local       LocalM   `json:"Local"`
	Minio       MinioM   `json:"Minio"`
}
type ConfigCA struct {
	PlatformUrl string    `json:"PlatformUrl"`
	KeyPair     KeyPairCA `json:"KeyPair"`
	Local       LocalCA   `json:"Local"`
}
type KeyPairCA struct {
	AutoConfig     bool   `json:"AutoConfig"`
	Algorithm      string `json:"Algorithm"`      //证书的生成算法rsa
	Bits           int    `json:"Bits"`           //证书生成的时候算法的位数有2096和1048
	PublicKeyPath  string `json:"PublicKeyPath"`  //私钥的存放路径
	PrivateKeyPath string `json:"PrivateKeyPath"` //公钥的存放路径
}

type LocalCA struct {
	Host        string `json:"Host"`
	Port        int    `json:"Port"`
	User        string `json:"User"`
	Password    string `json:"Password"`
	CurrentDir  string `json:"CurrentDir"`
	IDentity    string `json:"IDentity"`
	LoggerLevel string `json:"LoggerLevel"` //日志级别
	NoConsole   bool   `json:"NoConsole"`
}
