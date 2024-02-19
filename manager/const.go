package manager

type ManagerCmd string

const (
	//监管配置
	All                         ManagerCmd = "All"
	PlatformUrl                 ManagerCmd = "PlatformUrl"
	KeyPair                     ManagerCmd = "KeyPairAlgorithm"
	LocalHost                   ManagerCmd = "LocalHost"
	LocalPort                   ManagerCmd = "LocalPort"
	LocalUser                   ManagerCmd = "LocalUser"
	LocalPassword               ManagerCmd = "LocalPassword"
	LocalCurrentDir             ManagerCmd = "LocalCurrentDir"
	MinioEndPoint               ManagerCmd = "MinioEndPoint"
	MinioAccessKeyID            ManagerCmd = "MinioAccessKeyID"
	MinioSecretAccessKey        ManagerCmd = "MinioSecretAccessKey"
	MinioUseSSL                 ManagerCmd = "MinioUseSSL"
	MinioProductBucket          ManagerCmd = "MinioProductBucket"
	MinioProductDownload        ManagerCmd = "MinioProductDownload"
	MinioOfflineProductDownload ManagerCmd = "MinioOfflineProductDownload"

	//证书
	CertInput     ManagerCmd = "CertInput"
	CertInputRet  ManagerCmd = "CertInputRet"
	CertRemake    ManagerCmd = "CertRemake"
	CertRemakeRet ManagerCmd = "CertRemakeRet"
	CertSync      ManagerCmd = "CertSync"
	CertSyncRet   ManagerCmd = "CertSyncRet"
	CertShow      ManagerCmd = "CertShow"
	CertShowRet   ManagerCmd = "CertShowRet"
	CertOwner     ManagerCmd = "CertOwner"
	CertOwnerRet  ManagerCmd = "CertOwnerRet"

	//程序
	ExeGetLog       ManagerCmd = "ExeGetLog"
	ExeGetLogRet    ManagerCmd = "ExeGetLogRet"
	ExeGetStatus    ManagerCmd = "ExeGetStatus"
	ExeGetStatusRet ManagerCmd = "ExeGetStatusRet"
	ExeLocalPath    ManagerCmd = "ExeLocalPath"
	ExeLocalPathRet ManagerCmd = "ExeLocalPathRet"

	//数据产品
	DataGet        ManagerCmd = "DataGet"
	DataGetRet     ManagerCmd = "DataGetRet"
	DataControl    ManagerCmd = "DataControl"
	DataControlRet ManagerCmd = "DataControlRet"

	//版本
	Version    ManagerCmd = "Version"
	VersionRet ManagerCmd = "VersionRet"

	//算法
	AlgoRegister    ManagerCmd = "AlgoRegister"
	AlgoRegisterRet ManagerCmd = "AlgoRegisterRet"
	AlgoGetAll      ManagerCmd = "AlgoGetAll"
	AlgoGetAllRet   ManagerCmd = "AlgoGetAllRet"
	AlgoDelete      ManagerCmd = "AlgoDelete"
	AlgoDeleteRet   ManagerCmd = "AlgoDeleteRet"
)
