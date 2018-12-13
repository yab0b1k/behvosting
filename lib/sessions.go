package lib

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

func init() {
	config := session.ManagerConfig{
		CookieName:      "gohvosting",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}
	beego.GlobalSessions, _ = session.NewManager("memory", &config)
	go beego.GlobalSessions.GC()
}
