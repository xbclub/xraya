package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/xbclub/xraya/conf/report"
	_ "github.com/xbclub/xraya/pkg/plugin/pingtunnel"
	_ "github.com/xbclub/xraya/pkg/plugin/simpleobfs"
	_ "github.com/xbclub/xraya/pkg/plugin/socks5"
	_ "github.com/xbclub/xraya/pkg/plugin/ss"
	_ "github.com/xbclub/xraya/pkg/plugin/ssr"
	_ "github.com/xbclub/xraya/pkg/plugin/tcp"
	_ "github.com/xbclub/xraya/pkg/plugin/tls"
	_ "github.com/xbclub/xraya/pkg/plugin/trojanc"
	_ "github.com/xbclub/xraya/pkg/plugin/ws"
	"github.com/xbclub/xraya/pkg/util/log"
	"runtime"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	checkEnvironment()
	if runtime.GOOS == "linux" {
		checkTProxySupportability()
	}
	initConfigure()
	checkUpdate()
	hello()
	if err := run(); err != nil {
		log.Fatal("main: %v", err)
	}
}
