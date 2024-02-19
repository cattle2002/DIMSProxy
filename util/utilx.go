package util

import (
	"encoding/json"
	"os"
)

func ReadCopy() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	cfile, err := os.ReadFile("configc.json")
	if err != nil {
		panic(err)
	}
	var confCA ConfigCA
	var confM ConfigM
	err = json.Unmarshal(file, &confM)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(cfile, &confCA)
	if err != nil {
		panic(err)
	}

	if confCA.Local.User != "" && confCA.Local.Password != "" {
		if confM.Local.User == "" || confM.Local.Password == "" {
			confM.Local.User = confCA.Local.User
			confM.Local.Password = confCA.Local.Password
			//todo 序列化文件
			marshal, err := json.MarshalIndent(confM, "", " ")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile("config.json", marshal, 0666)
			if err != nil {
				panic(err)
			}
		} else {
			if confM.Local.User != confCA.Local.User || confM.Local.Password != confCA.Local.Password {
				panic("There is a difference in the user information between the two profiles，Please Update it")
			}
		}
	}
	if confM.Local.User != "" && confM.Local.Password != "" {
		if confCA.Local.User == "" || confCA.Local.Password == "" {
			confCA.Local.User = confM.Local.User
			confCA.Local.Password = confM.Local.Password
			//todo 序列化文件
			marshal, err := json.MarshalIndent(confCA, "", " ")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile("configc.json", marshal, 0666)
			if err != nil {
				panic(err)
			}
		} else {
			if confCA.Local.User != confM.Local.User || confCA.Local.Password != confM.Local.Password {
				panic("There is a difference in the user information between the two profiles，Please Update it")
			}
		}
	}
}
