package protocol

type AllReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
}

type AllRes struct {
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Config string `json:"Config"`
}

type PlatformUrlReq struct {
	Cmd         string `json:"Cmd"`
	Program     string `json:"Program"`
	PlatformUrl string `json:"PlatformUrl"`
}
type PlatformUrlRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type KeyPairReq struct {
	Cmd       string `json:"Cmd"`
	Program   string `json:"Program"`
	Algorithm string `json:"Algorithm"`
}
type KeyPairRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type LocalHostReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
	Host    string `json:"Host"`
}
type LocalHostRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type LocalPortReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
	Port    string `json:"Port"`
}
type LocalPortRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type LocalUserReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
	User    string `json:"User"`
}
type LocalUserRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type LocalPasswordReq struct {
	Cmd      string `json:"Cmd"`
	Program  string `json:"Program"`
	Password string `json:"Password"`
}
type LocalPasswordRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type LocalCaurrentDirReq struct {
	Cmd        string `json:"Cmd"`
	Program    string `json:"Program"`
	CurrentDir string `json:"CurrentDir"`
}
type LocalCauurentDirRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type MinioEndPointReq struct {
	Cmd      string `json:"Cmd"`
	Program  string `json:"Program"`
	EndPoint string `json:"EndPoint"`
}
type MinioEndPointRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type MinioAccessKeyIDReq struct {
	Cmd              string `json:"Cmd"`
	Program          string `json:"Program"`
	MinioAccessKeyID string `json:"MinioAccessKeyID"`
}
type MinioAccessKeyIDRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type MinioSecretAccessKeyReq struct {
	Cmd             string `json:"Cmd"`
	Program         string `json:"Program"`
	SecretAccessKey string `json:"SecretAccessKey"`
}
type MinioSecretAccessKeyRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type MinioUseSSLReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
	UseSSL  bool   `json:"UseSSL"`
}
type MinioUseSSlRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type MinioProductBucketReq struct {
	Cmd           string `json:"Cmd"`
	Program       string `json:"Program"`
	ProductBucket string `json:"ProductBucket"`
}
type MinioProductBucketRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type MinioProductDownloadReq struct {
	Cmd             string `json:"Cmd"`
	Program         string `json:"Program"`
	ProductDownload string `json:"ProductDownload"`
}
type MinioProductDownloadRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}
type MinioOfflineProductDownloadReq struct {
	Cmd                    string `json:"Cmd"`
	Program                string `json:"Program"`
	OfflineProductDownload string `json:"OfflineProductDownload"`
}
type MinioOfflineProductDownloadRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

type FindReq struct {
	Cmd string `json:"Cmd"`
	Ip  string `json:"Ip"`
}
type FindRes struct {
	Cmd string `json:"Cmd"`
	Ip  string `json:"Ip"`
}

// 导入证书
type CertInputReq struct {
	Cmd        string `json:"Cmd"`
	User       string `json:"User"`
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}
type CertInputRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
}

// 重新生成证书
type CertRemakeReq struct {
	Cmd string `json:"Cmd"`
}
type CertRemakeRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

// 同步平台证书
type CertSyncReq struct {
	Cmd string `json:"Cmd"`
}
type CertSyncRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

// 查看用户证书
type CertShowReq struct {
	Cmd string `json:"Cmd"`
}
type CertShowResPayload struct {
	User      string `json:"User"`
	PublicKey string `json:"PublicKey"`
}
type CertShowRes struct {
	IpAddr string               `json:"IpAddr"`
	Cmd    string               `json:"Cmd"`
	Code   int                  `json:"Code"`
	Msg    string               `json:"Msg"`
	Data   []CertShowResPayload `json:"Data"`
}
type CertOwnerReq struct {
	Cmd string `json:"Cmd"`
}
type CertOwnerRes struct {
	IpAddr     string `json:"IpAddr"`
	Cmd        string `json:"Cmd"`
	Code       int    `json:"Code"`
	Msg        string `json:"Msg"`
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}
type AlgoRegisterReq struct {
	Cmd      string `json:"Cmd"`
	AlgoJson string `json:"AlgoJson"`
}
type AlgoRegisterRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}
type AlgoGetAllReq struct {
	Cmd string `json:"Cmd"`
}
type AlgoGetAllRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

type AlgoDeleteReq struct {
	Cmd      string `json:"Cmd"`
	AlgoName string `json:"AlgoName"`
}
type AlgoDeleteRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}
type ExeGetLogReq struct {
	Cmd string `json:"Cmd"`
}
type ExeGetLogRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

type ExeGetStatusReq struct {
	Cmd string `json:"Cmd"`
}
type ExeGetStatusRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

type ExeLocalPathReq struct {
	Cmd string `json:"Cmd"`
}
type ExeLocalPathRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}
type DataGetReq struct {
	Cmd string `json:"Cmd"`
}
type DataPayload struct {
	ProductName string `json:"ProductName"`
}
type DataGetRes struct {
	IpAddr string        `json:"IpAddr"`
	Cmd    string        `json:"Cmd"`
	Code   int           `json:"Code"`
	Msg    string        `json:"Msg"`
	Data   []DataPayload `json:"Data"`
}

type DataControlReq struct {
	Cmd string `json:"Cmd"`
}

type DataControlRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}

type VersionReq struct {
	Cmd string `json:"Cmd"`
}
type VersionRes struct {
	IpAddr string `json:"IpAddr"`
	Cmd    string `json:"Cmd"`
	Code   int    `json:"Code"`
	Msg    string `json:"Msg"`
	Data   string `json:"Data"`
}
type AlgoJsonReq struct {
	AlgoJson string `json:"AlgoJson"`
}
