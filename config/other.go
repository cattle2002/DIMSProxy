package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"
)

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
func ethHost() string {
	dial, err := net.Dial("udp", "8.8.8.8:80") // Google的公共DNS服务器
	if err != nil {
		return "127.0.0.1"
	}
	addr := dial.LocalAddr().String()

	index := strings.LastIndex(addr, ":")
	return addr[:index]
}
func UpdateConfigCwd() {
	pwd := cwd()
	host := ethHost()

	var c Config

	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	c.Local.CurrentDir = pwd
	c.Local.EthHost = host
	marshal, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("config.json", marshal, 0666)
	if err != nil {
		panic(err)
	}
}
