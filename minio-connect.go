package main

// import "C"
// import (
// 	"context"
// 	"encoding/base64"
// 	"encoding/hex"
// 	"io"
// 	"os"
// 	"strings"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// 	"github.com/wumansgy/goEncrypt/aes"
// )

// // func NewMinioClient() { //初始化minio连接
// // 	mc, err := minio.New(config.Conf.Minio.EndPoint, &minio.Options{
// // 		Creds:  credentials.NewStaticV4(config.Conf.Minio.AccessKeyID, config.Conf.Minio.SecretAccessKey, ""),
// // 		Secure: config.Conf.Minio.UseSSL,
// // 	})
// // 	if err != nil {
// // 		log.Logger.Fatalf("connect minio error:%s", err.Error())
// // 	}
// // 	MinioClient = mc
// // }

// func NewMinioClient() { //初始化minio连接
// 	mc, err := minio.New("192.168.1.102:9000", &minio.Options{
// 		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
// 		Secure: false,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = mc.FPutObject(context.Background(), "test", "gg", "./main.go", minio.PutObjectOptions{})
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //export   SymmtricKeyDecryptPlus
// func SymmtricKeyDecryptPlus(data string, algo string, key string) (string, error) {
// 	decodeString, err := hex.DecodeString(key)
// 	if err != nil {
// 		return "", err
// 	}
// 	xdata, err := base64.StdEncoding.DecodeString(data)
// 	if err != nil {
// 		return "", err
// 	}
// 	plainText, err := aes.AesCbcDecrypt(xdata, decodeString, []byte("0000000000000000"))
// 	if err != nil {
// 		return "", nil
// 	}
// 	s := base64.StdEncoding.EncodeToString(plainText)
// 	return s, nil
// }

// //export SymmtricKeyDecrypt
// func SymmtricKeyDecrypt(data *C.char, algo *C.char, key *C.char) *C.char {
// 	gdata := C.GoString(data)
// 	gkey, err := hex.DecodeString(C.GoString(key))
// 	if err != nil {
// 		return C.CString("DIMSCASO-ERROR:" + err.Error())
// 	}
// 	xdata, err := base64.StdEncoding.DecodeString(gdata)
// 	if err != nil {
// 		return C.CString("DIMSCASO-ERROR:" + err.Error())
// 	}
// 	plainText, err := aes.AesCbcDecrypt(xdata, gkey, []byte("0000000000000000"))
// 	if err != nil {
// 		return C.CString("DIMSCASO-ERROR:" + err.Error())
// 	}
// 	s := base64.StdEncoding.EncodeToString(plainText)
// 	return C.CString(s)
// }

// func main() {
// 	b, err := os.ReadFile("file.enc")
// 	if err != nil {
// 		panic(err)
// 	}
// 	C.CString(string(b))
// 	c := SymmtricKeyDecrypt(C.CString(string(b)), C.CString("aes"), C.CString("cd46e0fa134dab432e748c37ad1be875"))
// 	b2, err2 := base64.StdEncoding.DecodeString(C.GoString(c))
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	// fmt.Println(string(b2))
// 	f, err3 := os.Create("dnc.txt")
// 	if err3 != nil {
// 		panic(err3)
// 	}
// 	r := strings.NewReader(string(b2))
// 	io.Copy(f, r)
// 	// b, err := hex.DecodeString("cd46e0fa134dab432e748c37ad1be875")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(b)
// 	// fmt.Println(len(b))

// 	// SymmtricKeyDecryptPluste()
// 	// NewMinioClient()
// }
