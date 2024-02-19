package handle

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
)

func VerifyUserName(buyer string) bool {
	log.Logger.Tracef("local user:%s,remote user:%s", config.Conf.Local.User, buyer)

	if config.Conf.Local.User == buyer {
		return true
	}
	return false
}
