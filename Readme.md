# 数据监管协同证书程序部署文档

## 环境要求

- Go 编程语言环境 go1.21 及以上
- SQLite 数据库

## 步骤

### 1. 克隆代码仓库

使用 Git 克隆项目代码到本地：

```bash
git clone https://github.com/cattle2002/DIMSProxy.git
cd DIMSProxy 
```

### 2.  安装依赖

```
1.go mod init DIMSProxy 
2.go mod tidy
```

### 3. 填写配置

config.json 是监管程序的配置

```
{
 "PlatformUrl": "ws://47.108.61.229:9081/api/websocket/1(这是你的必填项)",
 "KeyPair": {
  "AutoConfig": false,
  "Algorithm": "rsa",
  "Bits": 2048,
  "PublicKeyPath": "public.pub",
  "PrivateKeyPath": "private.key"
 },
 "Local": {
  "Host": "127.0.0.1",
  "Port": 5518,
  "CertPort": 5517,
  "User": "zcl01(这是你的必填项)",
  "Password": "zcl123456(这是你的必填项)",
  "CurrentDir": "",
  "Version": "Monitor@v20231203",
  "ManagerServicePort": 5520,
  "EthHost": "",
  "LoggerLevel": "trace",
  "NoConsole": false
 },
 "Minio": {
  "LifeDay": 1,
  "EndPoint": "47.108.20.64:9000(这是你的必填项)",
  "AccessKeyID": "g5fMoJm1xGdBSbgMrDpw(这是你的必填项)",
  "SecretAccessKey": "XdIi8sUQSqMR0k1XD1VDtpxOc0ECmWpkT6vnXkFI(这是你的必填项)",
  "UseSSL": false,
  "ProductBucket": "product",
  "ProductUpload": "productdownload",
  "OfflineProductDataEncrypt": "offlineproductdataencrypt"
 }
}

```

configc.json是证书程序的配置

```
{
 "PlatformUrl": "ws://47.108.61.229:9081/api/websocket/1(这是你的必填项)",
 "KeyPair": {
  "AutoConfig": false,
  "Algorithm": "rsa",
  "Bits": 2048,
  "PublicKeyPath": "public.pub",
  "PrivateKeyPath": "private.key"
 },
 "Local": {
  "Host": "127.0.0.1",
  "Port": 5517,
  "User": "zcl01(这是你的必填项)",
  "Password": "zcl123456(这是你的必填项)",
  "CurrentDir": "D:\\workdir\\DIMSProxy",
  "IDentity": "Master",
  "LoggerLevel": "trace",
  "NoConsole": false
 }
}
```

### 4.自行编译(libca.dll是仓库里面的DIMSCA编译得到的)

1.

```
编译监管程序
[windows]go build  -o server.exe  main.go (依赖当前目录下的libca.dll,当前lib/*作废)
[linux]go build  -o server  main.go(依赖当前目录下的libca.so,当前lib/*作废)
2.编译证书程序成为动态库
[windows] go  build -o  libca.dll -buildmode=c-shared  main.go
[linux] go  build -o  libca.so -buildmode=c-shared  main.go
3.提供给数据引擎(go & qt)的库
[windows] go  build -o  libca.dll -buildmode=c-shared  main_lib.go
[linux] go  build -o  libca.so -buildmode=c-shared  main_lib.go
```

### 5.使用编译好了的二进制程序

```
1.修改程序可执行权限:
[linux] chmod  +x 可执行程序.exe
2. 如何发现动态库不兼容当前glibc,请自行编译
3.最后执行可执行程序
[windows]./可执行程序.exe
[linux]./可执行程序
```

### 6.Thanks
